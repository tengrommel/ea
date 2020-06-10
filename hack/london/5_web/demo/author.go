package main

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var JWT_SECRET []byte = []byte("dshkfhdjk")

type Author struct {
	Id        string `json:"id, omitempty" validate:"omitempty, uuid"`
	FirstName string `json:"firstname, omitempty" validate:"required"`
	LastName  string `json:"lastname, omitempty" validate:"required"`
	UserName  string `json:"username, omitempty" validate:"required"`
	Password  string `json:"password, omitempty" validate:"required, gte=4"`
}

func RegisterEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var author Author
	json.NewDecoder(request.Body).Decode(&author)
	validate := validator.New()
	err := validate.Struct(author)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(author.Password), 10)
	author.Id = uuid.Must(uuid.NewV4(), nil).String()
	author.Password = string(hash)
	authors = append(authors, author)
	json.NewEncoder(response).Encode(authors)
}

func LoginEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var data Author
	json.NewDecoder(request.Body).Decode(&data)
	validate := validator.New()
	err := validate.StructExcept(data, "FirstName", "LastName")
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	for _, author := range authors {
		if author.UserName == data.UserName {
			err := bcrypt.CompareHashAndPassword([]byte(author.Password), []byte(data.Password))
			if err != nil {
				response.WriteHeader(500)
				response.Write([]byte(`{"message": "invalid password"}`))
				return
			}
			claims := CustomJWTClaim{
				Id: author.Id,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Local().Add(time.Hour).Unix(),
					Issuer:    "The Polyglot Developer",
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
			tokenString, _ := token.SignedString(JWT_SECRET)
			response.Write([]byte(`"token": "` + tokenString + `"}`))
			return
		}
	}
	response.Write([]byte(`{"message": "invalid password"}`))
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
	validate := validator.New()
	err := validate.StructExcept(changeItem, "FirstName", "LastName", "UserName", "Password")
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
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
				err = validate.Var(changeItem.Password, "gte=4")
				if err != nil {
					response.WriteHeader(500)
					response.Write([]byte(`{"message": "` + err.Error() + `"}`))
					return
				}
				hash, _ := bcrypt.GenerateFromPassword([]byte(changeItem.Password), 10)
				author.Password = string(hash)
			}
			authors[index] = author
			json.NewEncoder(response).Encode(authors)
			return
		}
	}
	json.NewEncoder(response).Encode(Author{})
}
