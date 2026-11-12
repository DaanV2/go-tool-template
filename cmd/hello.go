package cmd

import (
	"fmt"

	"charm.land/log/v2"
	"github.com/spf13/cobra"
)

func newHelloCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "hello [name]",
		Short:   "Print a greeting",
		Example: "go-tool-template hello world",
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := "world"
			if len(args) > 0 {
				name = args[0]
			}

			log.Debug("greeting", "name", name)
			_, err := fmt.Fprintf(cmd.OutOrStdout(), "Hello, %s!\n", name)

			return err
		},
	}
}
