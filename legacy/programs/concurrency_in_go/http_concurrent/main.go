package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN \n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d \n", url, resp.StatusCode)

		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			// file := strings.Split(url, "//")[1]
			file := generateRandomFileName()

			file += ".html"

			fmt.Printf("Writin response body to %s \n ", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	wg.Done()
}

func generateRandomFileName() (randString string) {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTZabcdefghijklmnopqrstuvwxyz1234567")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func main() {

	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com", "https://microsoft.com", "http://feeds.twit.tv/twit.xml", "http://feeds.bbci.co.uk/news/rss.xml"}

	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)

		fmt.Println(strings.Repeat("#", 20))
	}

	fmt.Println("No. of Gouritines: ", runtime.NumGoroutine())

	wg.Wait()

}
