package main

import (
	"os"

	"github.com/diptadas/golang-examples/cobra_example/cmd"
)

func main() {
	rootCmd := cmd.NewCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
