package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//forgot what api this belongs too. whoops.
const ArtistID = "4AK6F7OLvEQ5QYCBNiQWHq"

type Tracks struct {
	Track []Track `json:"tracks"`
}

type Track struct {
	Title string `json:"name"`
	Id    string `json:"id"`
}

//grabs songs from spotify and inserts them into redis
func getSongs() {
	//TODO: make request for top tracks to spotify
	url := "https://api.spotify.com/v1/artists/4AK6F7OLvEQ5QYCBNiQWHq/top-tracks?country=US"
	fmt.Println("about to make a request")
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("made request, resp code", res.StatusCode)
	defer res.Body.Close()
	topTracks := Tracks{}
	if err := json.NewDecoder(res.Body).Decode(&topTracks); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%#v\n", topTracks)
}
