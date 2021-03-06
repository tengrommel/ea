package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/couchbase/gocb.v1"
	"net/http"
)

var JWT_SECRET []byte = []byte("tehjkhajkf")

type GraphQLPayload struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type CustomJWTClaim struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func ValidateJWT(message string) (interface{}, error) {
	token, err := jwt.Parse(message, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return JWT_SECRET, nil
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
				var authors []Author
				query := gocb.NewN1qlQuery(`SELECT ` + bucket.Name() + `.* FROM ` + bucket.Name() + ` WHERE type = 'author'`)
				rows, err := bucket.ExecuteN1qlQuery(query, nil)
				if err != nil {
					return nil, err
				}
				var row Author
				for rows.Next(&row) {
					authors = append(authors, row)
				}
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
				//for _, author := range authors {
				//	if author.Id == id {
				//		return author, nil
				//	}
				//}
				//return nil, nil
				var author Author
				_, err := bucket.Get(id, &author)
				if err != nil {
					return nil, err
				}
				return author, nil
			},
		},
		"articles": &graphql.Field{
			Type: graphql.NewList(articleType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var articles []Article
				query := gocb.NewN1qlQuery(`SELECT ` + bucket.Name() + `.* FROM ` + bucket.Name() + ` WHERE  type = 'author'`)
				rows, err := bucket.ExecuteN1qlQuery(query, nil)
				if err != nil {
					return nil, err
				}
				var row Article
				for rows.Next(&row) {
					articles = append(articles, row)
				}
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
				var article Article
				bucket.Get(id, &article)
				_, err := bucket.Get(id, &article)
				if err != nil {
					return nil, err
				}
				return article, nil
				//for _, article := range articles {
				//	if article.Id == id {
				//		return article, nil
				//	}
				//}
				//return nil, nil
			},
		},
	},
})
var rootMotation *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"deleteAuthor": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["id"].(string)
				_, err := bucket.Remove(id, 0)
				if err != nil {
					return nil, err
				}
				return id, nil
				//for index, author := range authors {
				//	if author.Id == id {
				//		authors = append(authors[:index], authors[index+1:]...)
				//		return authors, nil
				//	}
				//}
				//return nil, nil
			},
		},
		"updateAuthor": &graphql.Field{
			Type: authorType,
			Args: graphql.FieldConfigArgument{
				"author": &graphql.ArgumentConfig{
					Type: authorInputType,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var changes Author
				mapstructure.Decode(params.Args["author"], &changes)
				validate := validator.New()
				mutation := bucket.MutateIn(changes.Id, 0, 0)
				if changes.FirstName != "" {
					mutation.Upsert("firstname", changes.FirstName, true)
				}
				if changes.LastName != "" {
					mutation.Upsert("lastname", changes.LastName, true)
				}
				if changes.UserName != "" {
					mutation.Upsert("username", changes.UserName, true)
				}
				if changes.Password != "" {
					err := validate.Var(changes.Password, "gte=4")
					if err != nil {
						return nil, err
					}
					hash, _ := bcrypt.GenerateFromPassword([]byte(changes.Password), 10)
					mutation.Upsert("password", string(hash), true)
				}
				mutation.Execute()
				return changes, nil
			},
		},
		"createArticle": &graphql.Field{
			Type: articleType,
			Args: graphql.FieldConfigArgument{
				"article": &graphql.ArgumentConfig{
					Type: articleInputType,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var article Article
				mapstructure.Decode(params.Args["article"], &article)
				decoded, err := ValidateJWT(params.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				validate := validator.New()
				err = validate.Struct(article)
				if err != nil {
					return nil, err
				}
				article.Id = uuid.Must(uuid.NewV4(), nil).String()
				article.Author = decoded.(CustomJWTClaim).Id
				article.Type = "article"
				bucket.Insert(article.Id, article, 0)
				//articles = append(articles, article)
				return article, nil
			},
		},
	},
})

var bucket *gocb.Bucket

func main() {
	fmt.Println("starting graphql application...")
	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "demo",
		Password: "abc123",
	})
	bucket, _ = cluster.OpenBucket("graphql", "")
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
			Context:        context.WithValue(context.Background(), "token", request.URL.Query().Get("token")),
		})
		json.NewEncoder(response).Encode(result)
	})
	methods := handlers.AllowedMethods(
		[]string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
		})
	headers := handlers.AllowedHeaders(
		[]string{
			"Content-Type",
			"Authorization",
			"X-Requested-With",
		})
	origins := handlers.AllowedOrigins(
		[]string{
			"*",
		})
	http.ListenAndServe(":12345", handlers.CORS(headers, methods, origins)(router))
}
