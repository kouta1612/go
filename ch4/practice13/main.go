package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type omdbSearchResult struct {
	Title    string
	Year     string
	Rated    string
	Released string
	Runtime  string
	Genre    string
	Director string
	Writer   string
	Actors   string
	Plot     string
	Language string
	Country  string
	Awards   string
	Poster   string
	Ratings  []struct {
		Source string
		Value  string
	}
	Metascore  string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}

func callOmdbAPI(url string) *omdbSearchResult {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	var result omdbSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "decode: %v\n", err)
		os.Exit(2)
	}
	resp.Body.Close()
	fmt.Println(result)
	return &result
}

func downloadPosterImage(imgURL string) {
	response := callOmdbAPI("https://omdbapi.com/?t=A+Beautiful+Mind")
	fmt.Println(response)
	// TODO API KEY取得してdownload処理追加
}

func main() {
	downloadPosterImage("a")
}
