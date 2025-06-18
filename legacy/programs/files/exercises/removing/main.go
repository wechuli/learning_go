package main

import (
	"log"
	"os"
)

func main() {

	filePath := "text.txt"
	err := os.Remove(filePath)

	if err != nil {
		log.Fatal(err)
	}
}
