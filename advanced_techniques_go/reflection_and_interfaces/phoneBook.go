package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"sort"
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

const CSVFILEPATH = "csv.data"

type PhoneBook []PhoneBookEntry

var data = PhoneBook{}
var index map[string]int

func PhoneBookPlay() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
	}

	err := setCSVFILE()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = readCSVFileNew(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()

	if err != nil {
		fmt.Println("Cannot create index")
		return
	}

	// Differenetiating between the commands

	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")
		if !matchTelNew(t) {
			fmt.Println("Not a valid telephone number: ", t)
			return
		}
		temp := initS(arguments[2], arguments[3], t)

		// If it was nil, there was an error

		if temp != nil {
			err := insert(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "delete":
		if len(arguments) != 3 {
			fmt.Println("Usage: delete Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTelNew(t) {
			fmt.Println("Not a valid telephone number: ", t)
			return
		}
		err := deleteEntry(t)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Number")
			return
		}

		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTelNew(t) {
			fmt.Println("Not a valid telephone number: ", t)
			return
		}
		temp := search(t)
		if temp == nil {
			fmt.Println("Number not found: ", t)
			return
		}
		fmt.Println(*temp)
	case "list":
		listContacts()
	default:
		fmt.Println("Not a valid option")

	}
}

func readCSVFileNew(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer f.Close()

	//CSV file read all at once

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

func saveCSVFileNew(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)

	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess}
		_ = csvwriter.Write(temp)
	}

	csvwriter.Flush()
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

// Initialized by the user - returns a pointer, if it returns nil, there was an error

func initS(N, S, T string) *PhoneBookEntry {
	if T == "" || S == "" {
		return nil
	}

	// Give LastAccess a value

	LastAccess := currentTime()
	return &PhoneBookEntry{Name: N, Surname: S, Tel: T, LastAccess: LastAccess}
}

func insert(pS *PhoneBookEntry) error {
	_, ok := index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}

	data = append(data, *pS)

	_ = createIndex()

	err := saveCSVFileNew(CSVFILE)
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

	// Update the index
	delete(index, key)

	err := saveCSVFileNew(CSVFILE)

	if err != nil {
		return err
	}

	return nil
}

func currentTime() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func search(key string) *PhoneBookEntry {
	i, ok := index[key]
	if !ok {
		return nil
	}

	data[i].LastAccess = currentTime()
	return &data[i]
}

func listContacts() {
	sort.Sort(PhoneBook(data))
	for _, v := range data {
		fmt.Println(v)
	}
}

func matchTelNew(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}

func setCSVFILE() error {
	filepath := os.Getenv("PHONEBOOK")
	if filepath != "" {
		CSVFILE = filepath
	}

	fileInfo, err := os.Stat(CSVFILE)
	if err != nil {
		fmt.Println("Creating ", CSVFILE)
		f, err := os.Create(CSVFILE)
		if err != nil {
			f.Close()
			return err
		}
		f.Close()
	}

	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		return fmt.Errorf("%s not a regular file", CSVFILE)
	}
	return nil

}

// Implement sort.Interface

func (a PhoneBook) Len() int {
	return len(a)
}

func (a PhoneBook) Less(i, j int) bool {
	if a[i].Surname == a[j].Surname {
		return a[i].Name < a[j].Name
	}

	return a[i].Surname < a[j].Surname
}

func (a PhoneBook) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
