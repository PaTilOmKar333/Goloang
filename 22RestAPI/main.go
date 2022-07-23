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
	foo(w)
}

func foo(w http.ResponseWriter) {
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

func TestPostarticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HomePage Endpoint Hit post")
}

func Handlerequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Homepage)
	myRouter.HandleFunc("/articles", Allarticles).Methods("GET")
	myRouter.HandleFunc("/articles", TestPostarticles).methods("POST")

	log.Fatal(http.ListenAndServe(":9065", myRouter))
}

func main() {
	Handlerequest()
}
