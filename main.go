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
	DB:       1,
})

func main() {
	start()
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/healthz", healthHandler)
	http.HandleFunc("/play", playHandler)
	http.ListenAndServe(":8080", nil)
}

func start() {
	// call service that gets and stores songs
	topTracks := getSongs()
	for _, track := range topTracks.Tracks {
		redisClient.Set(track.Title, track.Id, 0)
	}
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello world from Puffy!\n")
	fmt.Println("Puffy said hello")
}

func healthHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Puffy is OK!\n")
	fmt.Println("A health check was performed")
}

func playHandler(res http.ResponseWriter, req *http.Request) {
	song, _ := redisClient.RandomKey().Result()
	io.WriteString(res, "Play \""+song+"\" by One Direction!")
	fmt.Printf("\nPlay %s by One Direction for us!\n", song)

}
