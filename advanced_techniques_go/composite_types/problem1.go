package main

import "fmt"

func ExerciseOnePlay() {
	arr := [][]string{{"key1", "value1"}, {"key2", "value2"}, {"key3", "value3"}}
	mp := convertToMap(arr)
	fmt.Println(mp)
}

func convertToMap(arr [][]string) map[string]string {
	m := make(map[string]string)

	for _, entry := range arr {
		m[entry[0]] = entry[1]
	}

	return m
}
