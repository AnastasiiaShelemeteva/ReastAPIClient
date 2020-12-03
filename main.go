package main

import (
	"log"
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	showall, category, language := flags()

	if *showall {
		showCategories()
	} else {
		showJoke(category, language)
	}
}
