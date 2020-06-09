package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Author struct {
	Id        string `json:"id, omitempty"`
	FirstName string `json:"firstname, omitempty"`
	LastName  string `json:"lastname, omitempty"`
	UserName  string `json:"username, omitempty"`
	Password  string `json:"password, omitempty"`
}

func AuthorRetrieveAllEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	json.NewEncoder(response).Encode(authors)
}

func AuthorRetrieveEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	for _, author := range authors {
		if author.Id == params["id"] {
			json.NewEncoder(response).Encode(author)
			return
		}
	}
	json.NewEncoder(response).Encode(Author{})
}

func AuthorDeleteEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	for index, author := range authors {
		if author.Id == params["id"] {
			authors = append(authors[:index], authors[index+1:]...)
			json.NewEncoder(response).Encode(authors)
			return
		}
	}
	json.NewEncoder(response).Encode(Author{})
}

func AuthorUpdateEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	var changeItem Author
	json.NewDecoder(request.Body).Decode(&changeItem)
	for index, author := range authors {
		if author.Id == params["id"] {
			if changeItem.FirstName != "" {
				author.FirstName = changeItem.FirstName
			}
			if changeItem.LastName != "" {
				author.LastName = changeItem.LastName
			}
			if changeItem.UserName != "" {
				author.UserName = changeItem.UserName
			}
			if changeItem.Password != "" {
				author.Password = changeItem.Password
			}
			authors[index] = author
			json.NewEncoder(response).Encode(authors)
			return
		}
	}
	json.NewEncoder(response).Encode(Author{})
}
