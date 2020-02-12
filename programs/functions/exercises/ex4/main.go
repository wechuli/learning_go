package main

import (
	"fmt"
	"strings"
)

func searchItem(stringSlice []string, searchString string) (result bool) {

	// case sensitive search
	for _, value := range stringSlice {
		if value == searchString {
			return true
		}
	}

	//case insensitive search
	for _, value := range stringSlice {
		if strings.EqualFold(value, searchString) {
			return true
		}
	}

	return false
}

func main() {

	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "pig")
	fmt.Println(result) // => false

}
