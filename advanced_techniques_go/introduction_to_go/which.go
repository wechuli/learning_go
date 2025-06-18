package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Which() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}

	files := arguments[1:]

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, file := range files {
		searchExecutable(file, pathSplit)
	}
	
}

func searchExecutable(file string, pathSplit []string) {
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)

		//Does it exist

		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			if mode.IsRegular() {
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					// return
				}
			}
		}
	}

}
