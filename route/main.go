package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "hello world!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", Hello)

	log.Fatal(http.ListenAndServe("127.0.0.1:9999", router))
}
