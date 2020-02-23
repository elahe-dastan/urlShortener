package main

import (
	"urlShortener/config"
	"urlShortener/db"
	"urlShortener/service"
)

func main()  {
	//Give Configuration to db file
	constant := config.ReadConfig()
	db.SetConfig(constant)

	//This part of code generates all the random short ULRs possible and saves them into the database
	//This code executes only if the table containing the short URLs doesn't exist
	db.SaveShortURLs()

	//Creates a table for mapping each long URL to a short URL if not present in the database
	db.CreateMap()

	//Run post and get endpoints
	service.Run()
}
