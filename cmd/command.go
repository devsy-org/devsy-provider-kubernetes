package cmd

import (
	"context"
	"os"

	"github.com/devsy-org/devsy-provider-kubernetes/pkg/kubernetes"
	"github.com/devsy-org/devsy-provider-kubernetes/pkg/options"
	"github.com/devsy-org/devsy/pkg/driver"
	"github.com/spf13/cobra"
)

// CommandCmd holds the cmd flags.
type CommandCmd struct{}

// NewCommandCmd defines a command.
func NewCommandCmd() *cobra.Command {
	cmd := &CommandCmd{}
	commandCmd := &cobra.Command{
		Use:   "command",
		Short: "Command a container",
		RunE: func(_ *cobra.Command, args []string) error {
			options, err := options.FromEnv()
			if err != nil {
				return err
			}

			return cmd.Run(context.Background(), options)
		},
	}

	return commandCmd
}

// Run runs the command logic.
func (cmd *CommandCmd) Run(ctx context.Context, options *options.Options) error {
	return kubernetes.NewKubernetesDriver(options).CommandDevContainer(
		ctx,
		&driver.CommandParams{
			WorkspaceID: options.DevContainerID,
			User:        os.Getenv("DEVCONTAINER_USER"),
			Command:     os.Getenv("DEVCONTAINER_COMMAND"),
			Stdin:       os.Stdin,
			Stdout:      os.Stdout,
			Stderr:      os.Stderr,
		},
	)
}
