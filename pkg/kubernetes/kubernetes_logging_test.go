package kubernetes

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/devsy-org/devsy/pkg/log"
)

const (
	execCommand = "exec"
)

func TestKubectlCommandHelper(t *testing.T) {
	if os.Getenv("KUBECTL_TEST_HELPER") != "1" {
		return
	}
	for i, arg := range os.Args {
		if arg == "--" {
			if _, err := fmt.Fprint(os.Stdout, strings.Join(os.Args[i+1:], "|")); err != nil {
				t.Fatal(err)
			}
			os.Exit(0)
		}
	}
	os.Exit(2)
}

func TestKubectlCommandDoesNotLogArguments(t *testing.T) {
	logs := log.InitTestObserved(t, slog.LevelDebug)
	secret := "private-command-value-846297"
	driver := &KubernetesDriver{
		kubectl: os.Args[0],
	}
	args := []string{"-test.run=^TestKubectlCommandHelper$", "--", execCommand, secret}
	cmd := driver.buildCmd(context.Background(), args)
	cmd.Env = append(os.Environ(), "KUBECTL_TEST_HELPER=1")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		t.Fatalf("running kubectl command: %v", err)
	}
	if got := stdout.String(); got != "exec|"+secret {
		t.Fatalf("command output = %q", got)
	}

	driver = &KubernetesDriver{
		kubectl:    secret,
		namespace:  secret,
		kubeConfig: secret,
		context:    secret,
	}
	cmd = driver.buildCmd(context.Background(), []string{execCommand, secret})
	want := []string{
		secret,
		"--namespace",
		secret,
		"--kubeconfig",
		secret,
		"--context",
		secret,
		execCommand,
		secret,
	}
	if !reflect.DeepEqual(cmd.Args, want) {
		t.Fatalf("command arguments = %q, want %q", cmd.Args, want)
	}
	if logs.Len() != 2 {
		t.Fatalf("debug records = %d, want 2", logs.Len())
	}
	for _, entry := range logs.All() {
		if entry.Message != "Run kubectl command" ||
			strings.Contains(fmt.Sprint(entry.AttrsMap()), secret) {
			t.Fatalf("unexpected diagnostic: %s %v", entry.Message, entry.AttrsMap())
		}
	}
}
