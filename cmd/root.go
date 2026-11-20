package cmd

import (
	"errors"
	"os"
	"os/exec"

	"github.com/devsy-org/devsy/logging"
	"github.com/devsy-org/devsy/pkg/log"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

// NewRootCmd returns a new root command.
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "devsy-provider-kubernetes",
		Short:         "Kubernetes Provider commands",
		SilenceErrors: true,
		SilenceUsage:  true,

		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			config, invalid := logging.ResolveProvider(
				os.Getenv("DEVSY_LOG_LEVEL"),
				os.Getenv("DEVSY_LOG_FORMAT"),
			)
			if invalid.Level {
				return errors.New(
					"invalid DEVSY_LOG_LEVEL: expected error, warn, info, debug, or trace",
				)
			}
			if invalid.Format {
				return errors.New(
					"invalid DEVSY_LOG_FORMAT: expected auto, json, text, plain, or logfmt",
				)
			}
			log.Init(
				log.Config{
					Level:        logging.LevelName(config.Level),
					DefaultLevel: log.LevelInfoName,
					Format:       string(config.Format),
				},
			)
			return nil
		},
	}

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// build the root command
	rootCmd := BuildRoot()

	// execute command
	err := rootCmd.Execute()
	if err != nil {
		if exitErr, ok := err.(*ssh.ExitError); ok {
			os.Exit(exitErr.ExitStatus())
		}
		if exitErr, ok := err.(*exec.ExitError); ok {
			if len(exitErr.Stderr) > 0 {
				log.Error(string(exitErr.Stderr))
			}
			os.Exit(exitErr.ExitCode())
		}

		log.Fatal(err)
	}
}

// BuildRoot creates a new root command from the provider.
func BuildRoot() *cobra.Command {
	rootCmd := NewRootCmd()
	rootCmd.AddCommand(NewStartCmd())
	rootCmd.AddCommand(NewStopCmd())
	rootCmd.AddCommand(NewDeleteCmd())
	rootCmd.AddCommand(NewRunCmd())
	rootCmd.AddCommand(NewFindCmd())
	rootCmd.AddCommand(NewCommandCmd())
	rootCmd.AddCommand(NewTargetArchitectureCmd())
	return rootCmd
}
