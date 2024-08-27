package cmd

import (
	"fmt"
	"github.com/e-felix/sebas/internal/env"
	"github.com/spf13/cobra"
	"os"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read and print the environment variable defined in a `.env` file",
	Run: func(cmd *cobra.Command, args []string) {
		envs, err := env.ReadFile(args[0])

		if err != nil {
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		}

		for _, env := range envs {
			fmt.Printf("%v=%v\n", env.Key, env.Value)
		}
	},
}

func init() {
	EnvCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
