package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var authors []Author = []Author{
	{
		Id:        "author-1",
		FirstName: "Nic",
		LastName:  "Raboy",
		UserName:  "nraboy",
		Password:  "pass",
	},
	{
		Id:        "author-2",
		FirstName: "Maria",
		LastName:  "Raboy",
		UserName:  "Mraboy",
		Password:  "abc",
	},
}

var articles []Article = []Article{
	{
		Id:      "author-1",
		Author:  "author-1",
		Title:   "Blog Post 1",
		Content: " this is an example",
	},
}

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Write([]byte(`{"message": "Hello World"}`))
}

func main() {
	fmt.Println("Starting the application...")
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	router.HandleFunc("/authors", AuthorRetrieveAllEndpoint).Methods("GET")
	router.HandleFunc("/author/{id}", AuthorRetrieveEndpoint).Methods("GET")
	router.HandleFunc("/author/{id}", AuthorDeleteEndpoint).Methods("DELETE")
	router.HandleFunc("/author/{id}", AuthorUpdateEndpoint).Methods("PUT")
	router.HandleFunc("/articles", ArticleRetrieveAllEndpoint).Methods("GET")
	router.HandleFunc("/article/{id}", ArticleRetrieveEndpoint).Methods("GET")
	router.HandleFunc("/article/{id}", ArticleDeleteEndpoint).Methods("DELETE")
	router.HandleFunc("/article/{id}", ArticleUpdateEndpoint).Methods("PUT")
	router.HandleFunc("/article", ArticleCreateEndpoint).Methods("POST")
	http.ListenAndServe(":12345", router)
}
