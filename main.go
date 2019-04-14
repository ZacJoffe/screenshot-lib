package main

import (
	"./imagemagick"
	"log"
)

func main() {
	err := screenshot.Select()
	if err != nil {
		log.Fatal(err)
	}

	err = screenshot.Screen()
	if err != nil {
		log.Fatal(err)
	}
}
