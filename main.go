package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go behind Apache SS!")
}

func radio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Radio page from Go gateway")
}

func api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Api routing in development")
}

func main(){
	http.HandleFunc("/", home)
	http.HandleFunc("/radio", radio)
	http.HandleFunc("/api", api)

	log.Println("Go gateway listening on http://localhost:8080")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
