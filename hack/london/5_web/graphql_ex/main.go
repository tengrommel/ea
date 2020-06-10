package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/satori/go.uuid"
	"net/http"
)

type GraphQLPayload struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
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

var rootQuery *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"authors": &graphql.Field{
			Type: graphql.NewList(authorType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return authors, nil
			},
		},
		"author": &graphql.Field{
			Type: authorType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["id"].(string)
				for _, author := range authors {
					if author.Id == id {
						return author, nil
					}
				}
				return nil, nil
			},
		},
		"articles": &graphql.Field{
			Type: graphql.NewList(articleType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return articles, nil
			},
		},
		"article": &graphql.Field{
			Type: articleType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["id"].(string)
				for _, article := range articles {
					if article.Id == id {
						return article, nil
					}
				}
				return nil, nil
			},
		},
	},
})
var rootMotation *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"deleteAuthor": &graphql.Field{
			Type: graphql.NewList(authorType),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["id"].(string)
				for index, author := range authors {
					if author.Id == id {
						authors = append(authors[:index], authors[index+1:]...)
						return authors, nil
					}
				}
				return nil, nil
			},
		},
		"updateAuthor": &graphql.Field{
			Type: graphql.NewList(authorType),
			Args: graphql.FieldConfigArgument{
				"author": &graphql.ArgumentConfig{
					Type: authorInputType,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var changes Author
				mapstructure.Decode(params.Args["author"], &changes)
				for index, author := range authors {
					if author.Id == changes.Id {
						if changes.FirstName != "" {
							author.FirstName = changes.FirstName
						}
						if changes.LastName != "" {
							author.LastName = changes.LastName
						}
						if changes.FirstName != "" {
							author.FirstName = changes.FirstName
						}
						if changes.UserName != "" {
							author.UserName = changes.UserName
						}
						if changes.Password != "" {
							author.Password = changes.Password
						}
						authors[index] = author
						return authors, nil
					}
				}
				return nil, nil
			},
		},
		"createArticle": &graphql.Field{
			Type: graphql.NewList(articleType),
			Args: graphql.FieldConfigArgument{
				"article": &graphql.ArgumentConfig{
					Type: articleInputType,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var article Article
				mapstructure.Decode(params.Args["article"], &article)
				article.Id = uuid.Must(uuid.NewV4(), nil).String()
				article.Author = "teng"
				articles = append(articles, article)
				return articles, nil
			},
		},
	},
})

func main() {
	fmt.Println("starting graphql application...")
	router := mux.NewRouter()
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: rootQuery,
		})
	router.HandleFunc("/register", RegisterEndpoint).Methods("POST")
	router.HandleFunc("/login", LoginEndpoint).Methods("POST")
	router.HandleFunc("/graphql", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("content-type", "application/json")
		var payload GraphQLPayload
		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  payload.Query,
			VariableValues: payload.Variables,
		})
		json.NewEncoder(response).Encode(result)
	})
	http.ListenAndServe(":12345", router)
}
