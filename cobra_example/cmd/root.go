package cmd

import (
	"github.com/spf13/cobra"
)

var Verbose bool

func NewCmd() *cobra.Command {

	c := &cobra.Command{
		Use:     "echo",
		Short:   "Simple echo using Cobra",
		Example: "echo logo",
	}

	c.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	c.AddCommand(NewCmdLogo(), NewCmdEcho())
	return c
}

