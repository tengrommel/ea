package main

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

type Article struct {
	Id      string `json:"id, omitempty" validate:"omitempty, uuid"`
	Author  string `json:"author, omitempty" validate:"isdefault"`
	Title   string `json:"title, omitempty" validate:"required"`
	Content string `json:"content, omitempty" validate:"required"`
}

func ArticleCreateEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var article Article
	json.NewDecoder(request.Body).Decode(&article)
	token := context.Get(request, "decoded").(CustomJWTClaim)
	//tokenString := request.URL.Query().Get("token")
	//token, err := ValidateJWT(tokenString)
	//if err != nil {
	//	response.WriteHeader(500)
	//	response.Write([]byte(`{"message": "` + err.Error() + `"}`))
	//	return
	//}
	validate := validator.New()
	err := validate.Struct(article)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	article.Id = uuid.Must(uuid.NewV4(), nil).String()
	article.Author = token.Id
	articles = append(articles, article)
	json.NewEncoder(response).Encode(articles)
}

func ArticleRetrieveAllEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	json.NewEncoder(response).Encode(articles)
}

func ArticleRetrieveEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	for _, article := range articles {
		if article.Id == params["id"] {
			json.NewEncoder(response).Encode(article)
			return
		}
	}
	json.NewEncoder(response).Encode(Article{})
}

func ArticleDeleteEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	for index, article := range articles {
		if article.Id == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			json.NewEncoder(response).Encode(article)
			return
		}
	}
	json.NewEncoder(response).Encode(Article{})
}

func ArticleUpdateEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	var changeItem Article
	json.NewDecoder(request.Body).Decode(&changeItem)
	for index, article := range articles {
		if article.Id == params["id"] {
			if changeItem.Author != "" {
				article.Author = changeItem.Author
			}
			if changeItem.Content != "" {
				article.Content = changeItem.Content
			}
			if changeItem.Title != "" {
				article.Title = changeItem.Title
			}
			articles[index] = article
			json.NewEncoder(response).Encode(articles)
			return
		}
	}
	json.NewEncoder(response).Encode(Article{})
}
