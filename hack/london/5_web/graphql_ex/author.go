package main

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/graphql-go/graphql"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type Author struct {
	Id        string `json:"id, omitempty" validate:"omitempty, uuid"`
	FirstName string `json:"firstname, omitempty" validate:"required"`
	LastName  string `json:"lastname, omitempty" validate:"required"`
	UserName  string `json:"username, omitempty" validate:"required"`
	Password  string `json:"password, omitempty" validate:"required, gte=4"`
}

var authorType *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"firstname": &graphql.Field{
			Type: graphql.String,
		},
		"lastname": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var authorInputType *graphql.InputObject = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AuthorInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"firstname": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"lastname": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"username": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

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
	author.Id = uuid.Must(uuid.NewV4(), nil).String()
	hash, _ := bcrypt.GenerateFromPassword([]byte(author.Password), 10)
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
				response.Write([]byte(`{"message": "invalid password"`))
				return
			}
			claims := CustomJWTClaim{
				Id: author.Id,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Local().Add(time.Hour).Unix(),
					Issuer:    "The polyglot Developer",
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
			tokenString, _ := token.SignedString(JWT_SECRET)
			//json.NewEncoder(response).Encode(author)
			response.Write([]byte(`{"token": "` + tokenString + `"}`))
		}
	}
	json.NewEncoder(response).Encode(Author{})
}
