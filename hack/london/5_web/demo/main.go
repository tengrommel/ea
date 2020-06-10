package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type CustomJWTClaim struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

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

func ValidateJWT(message string) (interface{}, error) {
	token, err := jwt.Parse(message, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte("theployglotdeveloper"), nil
	})
	if err != nil {
		return nil, errors.New(`{"message": "` + err.Error() + `"}`)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var tokenData CustomJWTClaim
		mapstructure.Decode(claims, &tokenData)
		return tokenData, nil
	} else {
		return nil, errors.New(`{"message"` + err.Error() + `"}`)
	}
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
