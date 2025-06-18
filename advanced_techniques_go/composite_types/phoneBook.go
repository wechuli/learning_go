package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type PhoneBookEntry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

var CSVFILE = "/Users/wechuli/repos/learning_go/advanced_techniques_go/composite_types/csv.data"

var data = []PhoneBookEntry{}
var index map[string]int

var phoneBookIndex map[string]PhoneBookEntry

func PhoneBookMain() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}

	// If the CSVFILE does not exist, create an empty one

	_, err := os.Stat(CSVFILE)

	if err != nil {
		fmt.Println("Creating ", CSVFILE)

		f, err := os.Create(CSVFILE)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.Close()
	}

	fileInfo, err := os.Stat(CSVFILE)

	// Is it a regular file?

	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(CSVFILE, " not a regula file!")
		return
	}

	err = readPhoneBookCSVFile(CSVFILE)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()

	if err != nil {
		fmt.Println("Cannot create index")
		return
	}

	// Differenetiate between the commands

	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}

		t := strings.ReplaceAll(arguments[4], "-", "")
		if !matchTelNumber(t) {
			fmt.Println("Not a valid telephone number: ", t)
			return
		}

		temp := initPhoneBookS(arguments[2], arguments[3], t)
		if temp != nil {
			err := inserEntry(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Number")
			return
		}

		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTelNumber(t) {
			fmt.Println("Not a valid telephone number: ", t)
			return
		}
		temp := searchPhoneBook(t)
		if temp == nil {
			fmt.Println("Number not found: ", t)
			return
		}
		fmt.Println(*temp)
	case "list":
		listPhoneBookEntries()

	default:
		fmt.Println("Not a valie option")

	}
}

func readPhoneBookCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := PhoneBookEntry{
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: line[3],
		}

		data = append(data, temp)
	}

	return nil
}

func savePhoneBookCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer csvfile.Close()

	csvWriter := csv.NewWriter(csvfile)

	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess}
		_ = csvWriter.Write(temp)
	}

	csvWriter.Flush()
	return nil

}

func createIndex() error {
	index = make(map[string]int)
	for i, k := range data {
		key := k.Tel
		index[key] = i
	}
	return nil
}

func initPhoneBookS(N, S, T string) *PhoneBookEntry {
	if T == "" || S == "" {
		return nil
	}

	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &PhoneBookEntry{Name: N, Surname: S, Tel: T, LastAccess: LastAccess}
}

func inserEntry(pS *PhoneBookEntry) error {
	_, ok := index[(*pS).Tel]

	if ok {
		return fmt.Errorf("%s alread exists", pS.Tel)
	}

	data = append(data, *pS)

	// update index
	_ = createIndex()

	err := savePhoneBookCSVFile(CSVFILE)

	if err != nil {
		return err
	}

	return nil
}

func deleteEntry(key string) error {
	i, ok := index[key]

	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}

	data = append(data[:i], data[i+1:]...)

	delete(index, key)

	err := savePhoneBookCSVFile(CSVFILE)

	if err != nil {
		return err
	}

	return nil
}

func searchPhoneBook(key string) *PhoneBookEntry {
	i, ok := index[key]
	if !ok {
		return nil
	}

	data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	return &data[i]
}

func listPhoneBookEntries() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func matchTelNumber(s string) bool {
	re := regexp.MustCompile(`\d+$`)
	return re.MatchString(s)
}
