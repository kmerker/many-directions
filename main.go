package main

import (
	"fmt"
	"io"
	"net/http"
)

func indexHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello world from Puffy!\n")
	fmt.Println("Puffy said hello")
}

func healthHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Puffy is OK!\n")
	fmt.Println("A health check was performed")
}

func main() {
	//TODO: query spotify service for one direction information
	//TODO: insert returned info into postgres
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/healthz", healthHandler)
	http.ListenAndServe(":8080", nil)
}
