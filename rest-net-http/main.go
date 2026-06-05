package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func AllArticles(w http.ResponseWriter, r *http.Request) {
	articles:=Articles {
  	Article{ID:"1" ,Title:"Test Atricle",Desc:"test  description", Content:"Hello World"},
  	Article{ID:"2" ,Title:"Test Atricle",Desc:"test  description", Content:"Hello World"},
 	}

	fmt.Println("Endpoint hit: All article endpoint")
	json.NewEncoder(w).Encode(articles)
}

func SingleArticles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Print(params)

	articles:=Articles {
  	Article{ID:"1" ,Title:"Test Atricle",Desc:"test  description", Content:"Hello World"},
  	Article{ID:"2" ,Title:"Test Atricle",Desc:"test  description", Content:"Hello World"},
 	}

	for _, article := range articles {
		if article.ID == params["id"] {
			json.NewEncoder(w).Encode(article)
			return
		}
	}

	json.NewEncoder(w).Encode(&Articles{})
}

func PostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post endpoint hit")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// routes
	router.HandleFunc("/articles", AllArticles).Methods("GET")
	router.HandleFunc("/article/{id}", SingleArticles).Methods("GET")
	router.HandleFunc("/articles", PostArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
