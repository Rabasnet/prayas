package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	handleRequests()
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is my first api")
}

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/Articles", AllArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":8881", myRouter))
}

//Article is a struct with the name tags

type Article struct {
	Title string `json:"Title"`
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
}

//Articles is  slice of the Article
type Articles []Article

//AllArticles stores the actual data
func AllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "MR", Name: "Rbi Basnet", Age: 29},
		{Title: "MRS", Name: "tababi khadka", Age: 25},
		{Title: "MR", Name: "Roshan Basnet", Age: 25},
	}
	fmt.Printf("endpoint hit all articls on display")
	json.NewEncoder(w).Encode(articles)
}
