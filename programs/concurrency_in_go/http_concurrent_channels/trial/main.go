package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func checkAndSaveBody(url string, ch chan bool) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN \n", url)
		ch <- false
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d \n", url, resp.StatusCode)

		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]

			file += ".html"

			fmt.Printf("Writin response body to %s \n ", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}

			ch <- true
		}
	}
}

func main() {

	urls := []string{"https://gol65ang.org", "https://www.google.com", "https://www.medium.com", "https://microsoft.com"}

	// initialize the channel

	ch := make(chan bool)

	for _, url := range urls {

		go checkAndSaveBody(url, ch)

		f := <-ch

		fmt.Println(f)

		fmt.Println(strings.Repeat("#", 20))
	}

}
