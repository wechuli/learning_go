package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("info.txt")
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	file.Close()

}
