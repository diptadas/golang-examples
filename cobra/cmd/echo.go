package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func NewCmdEcho() *cobra.Command {
	var times int64

	c := &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long:  "Echo is for printing anything back to the screen.",
		Run: func(cmd *cobra.Command, args []string) {
			if verbose {
				fmt.Println("Info: Echo times:", times)
			}
			for times > 0 {
				fmt.Println("Echo: " + strings.Join(args, " "))
				times--
			}
		},
	}
	c.Flags().Int64VarP(&times, "times", "t", 1, "echo multiple times")
	c.AddCommand(NewCmdUpper())
	return c
}
