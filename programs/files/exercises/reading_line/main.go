package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// opening the file in read-only mode. The file must exist (in the current working directory)
	file, err := os.Open("poem.txt")

	// error handling
	if err != nil {
		log.Fatal(err)
	}

	// defer closing of the file
	defer file.Close()

	// the file value returned by os.Open() is wrapped in bufio.Scanner just like a bufferef reader
	scanner := bufio.NewScanner(file)

	// readin the whole file line by line
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
