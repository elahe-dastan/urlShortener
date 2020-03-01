package cmd

import (
	"fmt"
	"github.com/elahe-dastan/urlShortener_KGS/cmd/server"
	"github.com/elahe-dastan/urlShortener_KGS/cmd/setupdb"
	"github.com/elahe-dastan/urlShortener_KGS/config"
	"github.com/elahe-dastan/urlShortener_KGS/db"
	"github.com/elahe-dastan/urlShortener_KGS/middleware"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command {
	Use:   "urlShortener",
	Short: "Makes URLS shorter so they be can be memorized much easier",
	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	constant := config.ReadConfig()
	db.SetConfig(constant)
	middleware.SetConfig(constant)

	setupdb.Register(rootCmd)
	server.Register(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		exitFailure := 1
		os.Exit(exitFailure)
	}
}