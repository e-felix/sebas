package cmd

import (
	"os"

	envCli "github.com/e-felix/sebas/cli/cmd/env"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sebas [DOMAIN] [COMMAND] [ARG...]",
	Short: "Sebas is a general purpose tool to automate tasks per application project",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(envCli.EnvCmd)
}
