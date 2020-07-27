package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("web")))
	log.Fatal(http.ListenAndServe(":8080", router))
}
