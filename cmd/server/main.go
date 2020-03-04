package server

import (
	"github.com/elahe-dastan/urlShortener_KGS/config"
	"github.com/elahe-dastan/urlShortener_KGS/db"
	"github.com/elahe-dastan/urlShortener_KGS/service"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, constant config.Constants) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run server to serve the requests",
			Run: func(cmd *cobra.Command, args []string) {
				d := db.New(constant.DatabaseConfig)
				api := service.API{Map: struct{ DB *gorm.DB }{DB: d},
					ShortURL: struct{ DB *gorm.DB }{DB: d}}
				api.Run(constant.Log)
			},
		},
	)
}
