package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var logo = `
██████╗     ██╗    ██████╗     ████████╗     █████╗
██╔══██╗    ██║    ██╔══██╗    ╚══██╔══╝    ██╔══██╗
██║  ██║    ██║    ██████╔╝       ██║       ███████║
██║  ██║    ██║    ██╔═══╝        ██║       ██╔══██║
██████╔╝    ██║    ██║            ██║       ██║  ██║
╚═════╝     ╚═╝    ╚═╝            ╚═╝       ╚═╝  ╚═
`

func NewCmdLogo() *cobra.Command {
	c := &cobra.Command{
		Use:   "logo",
		Short: "Print the logo",
		Run: func(cmd *cobra.Command, args []string) {
			if verbose {
				fmt.Println("Info: Print logo")
			}
			fmt.Println(logo)
		},
	}
	return c
}
