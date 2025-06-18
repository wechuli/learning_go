package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	fmt.Println(os.Args[1:])

	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}
}
