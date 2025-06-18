package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// creating a new file

	var newFile *os.File
	fmt.Printf("%T \n", newFile)

	var err error

	newFile, err = os.Create("poem.txt")
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err)
	}

	// truncate, empty a file
	err = os.Truncate("poem.txt", 0)

	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err)
	}

	// close the file

	newFile.Close()

	file, err := os.OpenFile("poem.txt", os.O_APPEND, 0644)
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err)
	}

	file.Close()

	var fileInfo os.FileInfo

	fileInfo, err = os.Stat("poem.txt")
	p := fmt.Println

	p("File Name: ", fileInfo.Name())
	p("Size in bytes: ", fileInfo.Size())
	p("Last modified", fileInfo.ModTime())
	p("Is Directory", fileInfo.IsDir())
	p("Permissions: ", fileInfo.Mode())

	// non existent file
	fileInfo, err = os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err) {
			// log.Fatal("File does not exist")
		}
	}

	// moving/renaming file
	oldPath := "poem.txt"
	newPath := "epithet.txt"

	err = os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
		// log.Fatal(err)
	}

	// removing a file
	err = os.Remove("fake.txt")
	if err != nil {
		log.Fatal(err)
	}

}
