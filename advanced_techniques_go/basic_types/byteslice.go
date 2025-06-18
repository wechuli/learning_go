package main

import (
	"fmt"
)

func ByteSlicePlay() {
	b := make([]byte, 12)
	fmt.Println("Byte slice", b)

	b = []byte("Byte slice €")
	fmt.Println("Byte slice: ", b)
	fmt.Println("Length of byte slice = ", len(b))

	// Print byte slice contents as text
	fmt.Printf("Byte slice as text: %s\n", b)
	fmt.Println("Byte slice as text:", string(b))

	// Length of b
	fmt.Println("Length of b:", len(b))
}
