package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// scanner.Split(bufio.ScanWords)  configure how the scanner behaves

	success := scanner.Scan()

	if success == false {
		err = scanner.Err()
		if err == nil {
			log.Println("Scan was completed and it reached EOF")
		} else {
			log.Fatal(err)
		}
	}
	// fmt.Println("First line found: ", scanner.Text())

	// read whole string line by line

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}