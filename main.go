package main

import (
	"github.com/elahe-dastan/urlShortener_KGS/config"
	"github.com/elahe-dastan/urlShortener_KGS/db"
	"github.com/elahe-dastan/urlShortener_KGS/middleware"
	"github.com/elahe-dastan/urlShortener_KGS/service"
)

func main()  {
	//Give Configuration to db file
	constant := config.ReadConfig()
	db.SetConfig(constant)
	middleware.SetConfig(constant)

	//This part of code generates all the random short ULRs possible and saves them into the database
	//This code executes only if the table containing the short URLs doesn't exist
	db.SaveShortURLs()

	//Creates a table for mapping each long URL to a short URL if not present in the database
	db.CreateMap()

	//Run post and get endpoints
	service.Run()
}
