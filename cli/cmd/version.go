package cmd

import (
	"fmt"
	"github.com/e-felix/sebas/internal/util"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	_, b, _, _      = runtime.Caller(0)
	basepath        = filepath.Dir(b)
	versionFilename = "VERSION"
	versionFilePath = basepath + "/../../" + versionFilename
)

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		content, err := util.GetFileContent(versionFilename)

		if err != nil {
			fmt.Println("No version file found")
			os.Exit(1)
		}

		fmt.Printf("Sebas version: %s\n", content)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
