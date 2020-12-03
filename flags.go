package main

import (
	"flag"
)

func flags() (*bool, *string, *string) {

	showCategories := flag.Bool("showall", false, "display all categories")
	category := flag.String("cat", "Programming", "save option")
	language := flag.String("l", "en", "choose the language")
	flag.Parse()

	return showCategories, category, language

}
