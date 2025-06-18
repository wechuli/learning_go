package main

import (
	"fmt"
)

func Loops(){
// Traditional for loop

for i:=0; i<10; i++{
	fmt.Println(i*i," ")
}

j := 0

for ok := true;ok;ok=(j != 10){
fmt.Println(j*j, " ")
j++
}

aSlice := []int{-1,2,1,-1,2,-2}
for i,v := range aSlice{
	fmt.Println("index: ",i,"value: ",v)
}

fmt.Println()
}