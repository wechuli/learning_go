package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Create("info.txt")

	if err != nil {
		log.Fatal(err)
	}

	stringToWrite := "The Go gopher is an iconic mascot!"

	file.WriteString(stringToWrite)
	file.Close()

	fmt.Println(time.Now())
}
