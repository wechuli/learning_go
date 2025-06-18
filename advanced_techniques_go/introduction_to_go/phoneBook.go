package main

import "fmt"

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func InitializeData() {
	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471"})
	data = append(data, Entry{"Mary", "Doe", "2109416871"})
	data = append(data, Entry{"John", "Black", "2109416123"})
}

func Search(key string) *Entry {

	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func List() {
	for _, v := range data {
		fmt.Println(v)
	}
}
