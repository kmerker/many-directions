package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	redis "gopkg.in/redis.v3"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "", // no password set
	DB:       0,
})

func main() {
	start()
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/healthz", healthHandler)
	http.HandleFunc("/play", playHandler)
	http.HandleFunc("/ready", readyHandler)
	http.ListenAndServe(":8080", nil)
}

func start() {
	// call service that gets and stores songs
	topTracks := getSongs()
	time.Sleep(60 * time.Second)
	redisClient.Set("example song", "exampleid", 0)
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

func readyHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("\nA readiness check was performed at" + time.Now().String())
	song, err := redisClient.RandomKey().Result()
	if err != nil {
		fmt.Println("Looks like there was a problem getting stuff from redis")
		res.WriteHeader(http.StatusServiceUnavailable)
		io.WriteString(res, "Problem with redis command.")
	}
	if song == "" {
		fmt.Println("No songs in redis to find")
		res.WriteHeader(http.StatusServiceUnavailable)
		io.WriteString(res, "No songs!")
	}

	res.WriteHeader(http.StatusOK)
	io.WriteString(res, "JAM ON!")
	fmt.Println("JAM ON BRO")
}
