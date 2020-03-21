package server

import (
	"github.com/elahe-dastan/urlShortener/config"
	"github.com/elahe-dastan/urlShortener/db"
	"github.com/elahe-dastan/urlShortener/service"
	"github.com/elahe-dastan/urlShortener/store"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, cfg config.Config) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run server to serve the requests",
			Run: func(cmd *cobra.Command, args []string) {
				d := db.New(cfg.Database)
				api := service.API{
					Map:      store.NewMap(d),
					ShortURL: store.NewShortURL(d),
					Url:      cfg.URL}
				api.Run(cfg.Log)
			},
		},
	)
}
