package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// writing bytes to a file

	file, err := os.OpenFile(
		"poem.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	byteSlice := []byte("I learn Golang!")
	byteWritten, err := file.Write(byteSlice)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written: %d \n", byteWritten)

	// using ioutik

	bs := []byte("Go programming is cool!")
	err = ioutil.WriteFile("c.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}

	
}
