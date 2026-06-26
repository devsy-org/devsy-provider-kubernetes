package cmd

import (
	"context"

	"github.com/devsy-org/devsy-provider-kubernetes/pkg/kubernetes"
	"github.com/devsy-org/devsy-provider-kubernetes/pkg/options"
	"github.com/spf13/cobra"
)

// StopCmd holds the cmd flags.
type StopCmd struct{}

// NewStopCmd defines a command.
func NewStopCmd() *cobra.Command {
	cmd := &StopCmd{}
	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop a container",
		RunE: func(_ *cobra.Command, args []string) error {
			options, err := options.FromEnv()
			if err != nil {
				return err
			}

			return cmd.Run(context.Background(), options)
		},
	}

	return stopCmd
}

// Run runs the command logic.
func (cmd *StopCmd) Run(ctx context.Context, options *options.Options) error {
	return kubernetes.NewKubernetesDriver(options).
		StopDevContainer(ctx, options.DevContainerID)
}
