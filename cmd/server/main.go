package server

import (
	"github.com/elahe-dastan/urlShortener_KGS/service"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run server to serve the requests",
			Run: func(cmd *cobra.Command, args []string) {
				service.Run()
			},
		},
	)
}