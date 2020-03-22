package setupdb

import (
	"github.com/elahe-dastan/urlShortener/config"
	"github.com/elahe-dastan/urlShortener/db"
	"github.com/elahe-dastan/urlShortener/store"
	"github.com/spf13/cobra"
)

//type ShortURL struct {
//	Length int8
//}

func Register(root *cobra.Command, cfg config.Config) {
	c := cobra.Command{
		Use:   "setupdb",
		Short: "Manages database, creates and fills tables if don't exist",
		Run: func(cmd *cobra.Command, args []string) {
			d := db.New(cfg.Database)
			m := store.NewMap(d)
			l, _ := cmd.Flags().GetInt("length")
			u := store.ShortURL{
				DB:     d,
				Length: l,
			}

			//This part of code generates all the random short ULRs possible and saves them into the database
			//This code executes only if the table containing the short URLs doesn't exist
			u.Save()

			//Creates a table for mapping each long URL to a short URL if not present in the database
			m.Create()
		},
	}

	c.Flags().IntP("length", "l", 2, "KGS Length")

	root.AddCommand(
		&c,
	)
}
