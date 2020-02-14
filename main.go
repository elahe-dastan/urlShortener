package main

import (
	"urlShortener/configuration"
	"urlShortener/db"
	"urlShortener/services"
)

func main()  {
	//Give Configuration to db file
	constant := configuration.ReadConfig()
	db.SetConfiguration(constant)

	//This part of code generates all the random short ULRs possible and saves them into the database
	//This code executes only if the table containing the short URLs doesn't exist
	db.SaveRandomShortURLs()

	//Creates a table for mapping each long URL to a short URL if not present in the database
	db.CreateMapTable()

	//Run post and get endpoints
	services.RunServices()
}
