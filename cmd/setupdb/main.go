package setupdb

import (
	"github.com/elahe-dastan/urlShortener/config"
	"github.com/elahe-dastan/urlShortener/db"
	"github.com/elahe-dastan/urlShortener/store"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, cfg config.Config) {
	root.AddCommand(
		&cobra.Command{
			Use:   "setupdb",
			Short: "Manages database, creates and fills tables if don't exist",
			Run: func(cmd *cobra.Command, args []string) {
				d := db.New(cfg.Database)
				u := store.ShortURL{DB: d}
				m := store.Map{DB: d}

				//This part of code generates all the random short ULRs possible and saves them into the database
				//This code executes only if the table containing the short URLs doesn't exist
				u.Save()

				//Creates a table for mapping each long URL to a short URL if not present in the database
				m.Create()
			},
		},
	)
}
