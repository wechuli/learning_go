package main

import (
	"fmt"
	"os"

)

const (
	dbPort = 5432
)

func main() {

	host, user, password, dbname := os.Getenv("RSS_DB_HOST"), os.Getenv("RSS_DB_USERNAME"), os.Getenv("RSS_DB_PASSWORD"), os.Getenv("RSS_DB_NAME")

	fmt.Println("Starting program")
	app := App{}
	app.Initialize(host, dbPort, user, password, dbname)
	app.Run(":8080")

}

// $env:RSS_DB_NAME="rssfeeds"
// $env:RSS_DB_HOST="localhost"
// $env:RSS_DB_USERNAME="rssuser"
// $env:RSS_DB_PASSWORD="password"
