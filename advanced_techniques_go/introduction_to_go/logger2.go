package main

import (
	"fmt"
	"log"
	"os"
)

func Logger2() {
	if len(os.Args) != 2 {
		log.Fatal("Fatal: Hello World!")
	}
	log.Panic("Panic: Hello World!")

	fmt.Print("hi")
}
