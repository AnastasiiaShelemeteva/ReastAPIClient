package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//Category
type Category struct {
	Err        bool     `json:"error"`
	Categories []string `json:"categories"`
	Timestamp  int      `json:"timestamp"`
}

//Joke
type Jokes struct {
	Type     string `json:"type"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
	Joke     string `json:"joke"`
}

//Error
type Error struct {
	Err     bool     `json:"error"`
	Code    int      `json:"code"`
	Message []string `json:"causedBy"`
}

func showCategories() {
	response, err := http.Get("https://sv443.net/jokeapi/v2/categories")
	errCheck(err)
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	var categories Category
	json.Unmarshal([]byte(data), &categories)
	fmt.Println("Avaliable categories are:", strings.Join(categories.Categories, ", "))
}

func showJoke(category *string, language *string) {

	response, err := http.Get("https://sv443.net/jokeapi/v2/joke/" + *category + "?format=json&lang=" + *language + "&blacklistFlags=sexist,racist&amount=1")
	errCheck(err)
	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	apiErrCheck(data)

	var joke Jokes
	json.Unmarshal([]byte(data), &joke)
	if joke.Type == "single" {
		fmt.Println(joke.Joke)
	} else if joke.Type == "twopart" {
		fmt.Printf("%s\n%s\n", joke.Setup, joke.Delivery)
	}

}

func apiErrCheck(data []byte) {
	var err Error
	json.Unmarshal([]byte(data), &err)
	if err.Err {
		log.Fatalf("Error code: %d\nError message: %s", err.Code, strings.Join(err.Message, ""))
	}
}
