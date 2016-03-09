package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//forgot what api this belongs too. whoops.
const ArtistID = "4AK6F7OLvEQ5QYCBNiQWHq"

type Tracks struct {
	Track []Track
}

type Track struct {
	Title string `json:"name"`
	Id    string `json:"id"`
}

//grabs songs from spotify and inserts them into redis
func getSongs() {
	//TODO: make request for top tracks to spotify
	url := "https://api.spotify.com/v1/artists/4AK6F7OLvEQ5QYCBNiQWHq/to-tracks?country=US"
	fmt.Println("about to make a request")
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println("serializing...")
	topTracks := Tracks{}
	if err := json.Unmarshal(body, &topTracks); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v", topTracks)
}
