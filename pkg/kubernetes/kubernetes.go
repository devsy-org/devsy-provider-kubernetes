package kubernetes

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/devsy-org/devsy-provider-kubernetes/pkg/options"
	"github.com/devsy-org/devsy/pkg/command"
	"github.com/devsy-org/devsy/pkg/driver"
	"github.com/devsy-org/devsy/pkg/log"
	corev1 "k8s.io/api/core/v1"
)

func NewKubernetesDriver(options *options.Options) driver.Driver {
	kubectl := "kubectl"
	if options.KubectlPath != "" {
		kubectl = options.KubectlPath
	}

	if options.KubernetesNamespace != "" {
		log.Debugf("Use Kubernetes Namespace '%s'", options.KubernetesNamespace)
	}
	if options.KubernetesConfig != "" {
		log.Debugf("Use Kubernetes Config '%s'", options.KubernetesConfig)
	}
	if options.KubernetesContext != "" {
		log.Debugf("Use Kubernetes Context '%s'", options.KubernetesContext)
	}
	return &KubernetesDriver{
		kubectl: kubectl,

		kubeConfig: options.KubernetesConfig,
		context:    options.KubernetesContext,
		namespace:  options.KubernetesNamespace,

		options: options,
	}
}

type KubernetesDriver struct {
	kubectl string

	kubeConfig string
	namespace  string
	context    string

	options *options.Options
}

func (k *KubernetesDriver) StopDevContainer(ctx context.Context, workspaceId string) error {
	workspaceId = getID(workspaceId)

	// delete pod
	out, err := k.buildCmd(ctx, []string{"delete", "po", workspaceId, "--ignore-not-found"}).
		CombinedOutput()
	if err != nil {
		return fmt.Errorf("delete pod: %s: %w", string(out), err)
	}

	return nil
}

func (k *KubernetesDriver) DeleteDevContainer(ctx context.Context, workspaceId string) error {
	workspaceId = getID(workspaceId)

	// delete pod
	log.Infof("Delete pod '%s'...", workspaceId)
	err := k.deletePod(ctx, workspaceId)
	if err != nil {
		return err
	}

	// delete pvc
	log.Infof("Delete persistent volume claim '%s'...", workspaceId)
	out, err := k.buildCmd(ctx, []string{"delete", "pvc", workspaceId, "--ignore-not-found", "--grace-period=5"}).
		CombinedOutput()
	if err != nil {
		return fmt.Errorf("delete pvc: %s: %w", string(out), err)
	}

	// delete role binding & service account
	if k.options.ClusterRole != "" {
		log.Infof("Delete role binding '%s'...", workspaceId)
		out, err := k.buildCmd(ctx, []string{"delete", "rolebinding", workspaceId, "--ignore-not-found"}).
			CombinedOutput()
		if err != nil {
			return fmt.Errorf("delete role binding: %s: %w", string(out), err)
		}
	}

	// delete pull secret
	if k.options.KubernetesPullSecretsEnabled != "" {
		log.Infof("Delete pull secret '%s'...", workspaceId)
		err := k.DeletePullSecret(ctx, getPullSecretsName(workspaceId))
		if err != nil {
			return err
		}
	}

	return nil
}

func (k *KubernetesDriver) CommandDevContainer(
	ctx context.Context,
	params *driver.CommandParams,
) error {
	workspaceId := getID(params.WorkspaceID)

	args := []string{"exec", "-c", "devsy"}
	if params.Stdin != nil {
		args = append(args, "-i")
	}
	args = append(args, workspaceId)
	if params.User != "" && params.User != "root" {
		args = append(args, "--", "su", params.User, "-c", params.Command)
	} else {
		args = append(args, "--", "sh", "-c", params.Command)
	}

	return k.runCommand(
		ctx,
		args,
		cmdIO{stdin: params.Stdin, stdout: params.Stdout, stderr: params.Stderr},
	)
}

func (k *KubernetesDriver) GetDevContainerLogs(
	ctx context.Context,
	workspaceID string,
	stdout io.Writer,
	stderr io.Writer,
) error {
	workspaceID = getID(workspaceID)

	args := []string{"logs", "pods/" + workspaceID, "-c", "devsy"}

	return k.runCommand(ctx, args, cmdIO{stdout: stdout, stderr: stderr})
}

func (k *KubernetesDriver) deletePod(ctx context.Context, podName string) error {
	out, err := k.buildCmd(ctx, []string{"delete", "po", podName, "--ignore-not-found", "--grace-period=10"}).
		CombinedOutput()
	if err != nil {
		return fmt.Errorf("delete pod: %s: %w", string(out), err)
	}

	return nil
}

func (k *KubernetesDriver) getDevContainerPvc(
	ctx context.Context,
	id string,
) (*corev1.PersistentVolumeClaim, *DevContainerInfo, error) {
	// try to find pvc
	out, err := k.buildCmd(
		ctx, []string{"get", "pvc", id, "--ignore-not-found", "-o", "json"},
	).Output()
	if err != nil {
		return nil, nil, command.WrapCommandError(out, err)
	} else if len(out) == 0 {
		return nil, nil, nil
	}

	// try to unmarshal pvc
	pvc := &corev1.PersistentVolumeClaim{}
	err = json.Unmarshal(out, pvc)
	if err != nil {
		return nil, nil, fmt.Errorf("unmarshal pvc: %w", err)
	} else if pvc.Annotations == nil ||
		pvc.Annotations[DevPodInfoAnnotation] == "" {
		return nil, nil, fmt.Errorf(
			"pvc is missing dev container info annotation",
		)
	}

	// get container info
	containerInfo := &DevContainerInfo{}
	err = json.Unmarshal(
		[]byte(pvc.GetAnnotations()[DevPodInfoAnnotation]),
		containerInfo,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("decode dev container info: %w", err)
	}

	return pvc, containerInfo, nil
}

func (k *KubernetesDriver) buildCmd(ctx context.Context, args []string) *exec.Cmd {
	newArgs := []string{}
	if k.namespace != "" {
		newArgs = append(newArgs, "--namespace", k.namespace)
	}
	if k.kubeConfig != "" {
		newArgs = append(newArgs, "--kubeconfig", k.kubeConfig)
	}
	if k.context != "" {
		newArgs = append(newArgs, "--context", k.context)
	}
	newArgs = append(newArgs, args...)
	log.Debugf("Run command: %s %s", k.kubectl, strings.Join(newArgs, " "))
	//nolint:gosec // kubectl path is from trusted configuration
	return exec.CommandContext(ctx, k.kubectl, newArgs...)
}

type cmdIO struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

func (k *KubernetesDriver) runCommand(
	ctx context.Context,
	args []string,
	io cmdIO,
) error {
	return k.runCommandInDir(ctx, "", args, io)
}

func (k *KubernetesDriver) runCommandInDir(
	ctx context.Context,
	dir string,
	args []string,
	io cmdIO,
) error {
	cmd := k.buildCmd(ctx, args)
	cmd.Dir = dir
	cmd.Stdin = io.stdin
	cmd.Stdout = io.stdout
	cmd.Stderr = io.stderr
	return cmd.Run()
}
