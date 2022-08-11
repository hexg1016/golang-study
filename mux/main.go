package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!\n")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", Hello)

	log.Fatal(http.ListenAndServe("127.0.0.1:9999", router))
}
