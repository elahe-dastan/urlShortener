package cmd

import (
	"fmt"
	"os"

	"github.com/elahe-dastan/urlShortener_KGS/cmd/server"
	"github.com/elahe-dastan/urlShortener_KGS/cmd/setupdb"
	"github.com/elahe-dastan/urlShortener_KGS/config"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "urlShortener",
		Short: "Makes URLS shorter so they be can be memorized much easier",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	exitFailure := 1

	constant := config.ReadConfig()

	setupdb.Register(rootCmd, constant)
	server.Register(rootCmd, constant)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(exitFailure)
	}
}
