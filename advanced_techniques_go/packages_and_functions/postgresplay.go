package main

import (
	"fmt"
	"math/rand"
	"packages/postgres"
	"time"
)

var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func PostgresPlay() {
	postgres.Hostname = "localhost"
	postgres.Port = 6004
	postgres.Username = "educative"
	postgres.Password = "pass"
	postgres.Database = "master"

	data, err := postgres.ListUsers()

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range data {
		fmt.Println(v)
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)
	random_username := getString(10)

	t := postgres.Userdata{
		Username:    random_username,
		Name:        "Paul",
		Surname:     "Wechuli",
		Description: "This is me"}

	id := postgres.AddUser(t)

	if id == -1 {
		fmt.Println("There was an error adding the user", t.Username)
	}

}

func getString(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return temp
}
