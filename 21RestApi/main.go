package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Name    string `json:"firstName"`
	Surname string `json:"lastName"`
	Age     int    `json:"age"`
}

type Articles []Article

func Allarticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Name: "omkar", Surname: "patil", Age: 27},
		Article{Name: "saurabh", Surname: "patil", Age: 97},
	}
	fmt.Println("All articles endpoint")

	json.NewEncoder(w).Encode(articles)
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HomePage Endpoint Hit")
}

func Handlerequest() {
	r := mux.NewRouter()
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/articles", Allarticles)

	log.Fatal(http.ListenAndServe(":9065", r))
}

func main() {
	Handlerequest()
}
