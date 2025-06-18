package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filePath := "poem.txt"

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// ioutil.ReadAll() reads from the file until an error or EOF and returns the data it read

	fileContents, err := ioutil.ReadAll(file)

	if err != nil {
		log.
			Fatal(err)
	}

	fmt.Printf("Data as string: \n %s \n", fileContents)

}
