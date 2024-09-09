package cmd

import (
	"fmt"
	"github.com/e-felix/sebas/internal/database"
	"github.com/e-felix/sebas/internal/env"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var saveCmd = &cobra.Command{
	Use:   "save [FILEPATH] [PROJECT]",
	Short: "Save the environment variables defined in a `.env` file for a project",
	Run: func(cmd *cobra.Command, args []string) {
		envs, err := env.ReadFile(args[0])

		if err != nil {
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		}

		for _, env := range envs {
			fmt.Printf("%v=%v\n", env.Key, env.Value)
		}

		db, err := database.Connect("./sqlite.db")
		defer db.Close()

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Connected to database")
	},
}

func init() {
	EnvCmd.AddCommand(saveCmd)
}
