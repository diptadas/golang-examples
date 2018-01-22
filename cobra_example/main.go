package main

import (
	"golang-examples/cobra_example/cmd"
	"os"
)

func main() {
	rootCmd := cmd.NewCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
