package main

import (
	"fmt"
	// "os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rssuser"
	password = "password"
	dbname   = "rssfeeds"
)

func main() {

	// host, port, user, password, dbname := os.Getenv("RSS_DB_HOST"), 5432, os.Getenv("RSS_DB_USERNAME"), os.Getenv("RSS_DB_PASSWORD"), os.Getenv("RSS_DB_NAME")

	fmt.Println("Starting program")
	app := App{}
	app.Initialize(host, port, user, password, dbname)
	app.Run(":8080")

}
