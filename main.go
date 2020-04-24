package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("got request for /")
	fmt.Fprint(w, "Welcome!\n")
}

func hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "hello, %s!\n", vars["name"])
}

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi, %s!\n", r.URL.Query().Get("name"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/hello/{name}", hello).Methods("POST")
	r.HandleFunc("/hi", hi).Queries("name", "scguo")

	if err := http.ListenAndServe(":18088", r); err != nil {
		fmt.Printf("Start error %v", err)
	}
}
