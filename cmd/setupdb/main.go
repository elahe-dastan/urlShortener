package setupdb

import (
	"github.com/elahe-dastan/urlShortener_KGS/db"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "setupdb",
			Short: "Manages database, creates and fills tables if don't exist",
			Run: func(cmd *cobra.Command, args []string) {
				//This part of code generates all the random short ULRs possible and saves them into the database
				//This code executes only if the table containing the short URLs doesn't exist
				db.SaveShortURLs()
				//Creates a table for mapping each long URL to a short URL if not present in the database
				db.CreateMap()
			},
		},
	)
}