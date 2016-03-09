package main

import (
	"fmt"
	"io"
	"net/http"

	redis "gopkg.in/redis.v3"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {
	// call start func which does the below
	//TODO: query spotify service for one direction information
	//TODO: insert returned info into redis

	start()
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/healthz", healthHandler)
	fmt.Println("serving on port 8080")
	http.ListenAndServe(":8080", nil)
}

func start() {
	// call service that gets and stores songs
	getSongs()
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello world from Puffy!\n")
	fmt.Println("Puffy said hello")
}

func healthHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Puffy is OK!\n")
	fmt.Println("A health check was performed")
}
