package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"_id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// let's declare a global Article array
// that we can then populate in our main function
// to simulate a database

var Articles []Article

// return all the articles
func returnAllArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(Articles)

}

// return single article
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {

	var articleNotFoundError = ErrorResponse{Success: false, Message: "Article unavailable"}
	articleFound := false

	vars := mux.Vars(r)
	key := vars["id"]
	// loop over all of the articles, if the article id equals the key we pass in return the article encoded as JSON
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
			articleFound = true
		}
	}

	if !articleFound {
		json.NewEncoder(w).Encode(articleNotFoundError)
	}

}

// create a new article
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	//  get the body of our POST request
	// return the string response containing the request body

	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article

	json.Unmarshal(reqBody, &article)

	// update our global Articles array to include the new array

	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

// delete an article

func deleteArticle(w http.ResponseWriter, r *http.Request){
	
}

// homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	//  creates a new instance of a mux router

	myRouter := mux.NewRouter().StrictSlash(true)

	// replace http.HandleFunc with myRouter.HandleFunc

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)

	// NOTE: Ordering is important here! This has to be defined before the other `/article` endpoint.
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")

	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	// finally, instead of passing in nil, we want to pass in our newly created router as the second argument

	log.Fatal(http.ListenAndServe(":10000", myRouter))

}

func main() {

	fmt.Println("REST API v2.0 - Mux Routers")

	Articles = []Article{
		Article{Id: "1", Title: "Test Title", Desc: "Test Description", Content: "Test Content - Hello world ya'll"},
		Article{Id: "2", Title: "Test Title2", Desc: "Test Description2", Content: "Test Content - Hello world ya'll 2"},
		Article{Id: "3", Title: "Test Title3", Desc: "Test Description3", Content: "Test Content - Hello world ya'll 3"},
		Article{Id: "4", Title: "Test Title4", Desc: "Test Description4", Content: "Test Content - Hello world ya'll 4"},
	}

	handleRequests()
}
