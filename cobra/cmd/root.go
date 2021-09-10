package cmd

import (
	"github.com/spf13/cobra"
)

var verbose bool

func NewCmd() *cobra.Command {
	c := &cobra.Command{
		Use:     "my-cli",
		Short:   "Simple CLI using Cobra",
		Example: "my-cli logo",
	}
	c.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	c.AddCommand(NewCmdLogo(), NewCmdEcho())
	return c
}
