package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func NewCmdUpper() *cobra.Command {
	c := &cobra.Command{
		Use:   "upper [string to echo]",
		Short: "Echo anything to the screen in upper case",
		Run: func(cmd *cobra.Command, args []string) {
			if verbose {
				fmt.Println("Info: Echo in upper case")
			}
			fmt.Println("Echo Upper: " + strings.ToUpper(strings.Join(args, " ")))
		},
	}
	return c
}
