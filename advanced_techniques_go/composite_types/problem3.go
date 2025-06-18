package main

import (
	"fmt"
	"os"
)

type Arg struct {
	Index int
	Value string
}

func ExerciseThreePlay() {
	args := make([]Arg, len(os.Args))

	for index, arg := range os.Args {
		argStruct := Arg{index, arg}
		args[index] = argStruct
	}
	for _, arg := range args {
		fmt.Printf("Arg[%d]: %s\n", arg.Index, arg.Value)
	}
}
