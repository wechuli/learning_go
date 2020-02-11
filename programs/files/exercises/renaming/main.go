package main

import (
	"log"
	"os"
)

func main() {

	filePath := "info.txt"
	// checking if file exists
	_, err := os.Stat(filePath)
	//error handling

	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist!")
		}
	}

	// renaming the file
	err = os.Rename(filePath, "data.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}
