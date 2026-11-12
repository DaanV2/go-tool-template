// Package cmd contains the CLI commands.
package cmd

import (
	"context"

	"charm.land/fang/v2"
	"github.com/spf13/cobra"

	"github.com/DaanV2/go-tool-template/infrastructure/logging"
	"github.com/DaanV2/go-tool-template/infrastructure/version"
)

// NewRootCmd builds the root command and all its subcommands.
func NewRootCmd() *cobra.Command {
	logOpts := logging.Options{}

	root := &cobra.Command{
		Use:   "go-tool-template",
		Short: "A template for Go CLI tools",
		Long:  "go-tool-template is a starting point for Go CLI tools built with cobra, fang and charm log.",
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			return logging.Setup(logOpts)
		},
		SilenceUsage: true,
	}

	root.PersistentFlags().StringVar(&logOpts.Level, "log-level", "info", "log level (debug, info, warn, error, fatal)")
	root.PersistentFlags().StringVar(&logOpts.Format, "log-format", "text", "log format (text, json, logfmt)")

	root.AddCommand(newHelloCmd())

	return root
}

// Execute runs the root command with fang.
func Execute(ctx context.Context) error {
	opts := []fang.Option{fang.WithVersion(version.Version)}
	if version.Commit != "" {
		opts = append(opts, fang.WithCommit(version.Commit))
	}

	return fang.Execute(ctx, NewRootCmd(), opts...)
}
