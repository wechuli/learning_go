package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	Time        string `json:"time"`
	Information Info   `json:"info"`
}

type Info struct {
	Description string  `json:"desc"`
	Value       float64 `json:"value"`
}

func main() {
	fmt.Println("Hi there")

	jsonString := `
	{
	"name":"漢batteryʆ",
	"capacity":40,
	"time":"2019-01-21T19:07:28Z",
	"info":{
		"desc":"a sensor reading",
		"value":85.95
	 }
	}
	`

	var reading SensorReading

	err := json.Unmarshal([]byte(jsonString), &reading)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", reading)

	var reading2 map[string]interface{}

	err2 := json.Unmarshal([]byte(jsonString), &reading2)

	if err2 != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", reading2)

}
