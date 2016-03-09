package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//forgot what api this belongs too. whoops.
const ArtistID = "4AK6F7OLvEQ5QYCBNiQWHq"

type TopTracks struct {
	Tracks []Track `json:"tracks"`
}

type Track struct {
	Title string `json:"name"`
	Id    string `json:"id"`
}

//grabs songs from spotify and inserts them into redis
func getSongs() TopTracks {
	url := "https://api.spotify.com/v1/artists/4AK6F7OLvEQ5QYCBNiQWHq/top-tracks?country=US"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	topTracks := TopTracks{}
	if err := json.NewDecoder(res.Body).Decode(&topTracks); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return topTracks
}
