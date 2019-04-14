package main

import (
	"./screenshot"
	"log"
)

func main() {
	_, err := screenshot.Select()
	if err != nil {
		log.Fatal(err)
	}

	_, err = screenshot.Screen()
	if err != nil {
		log.Fatal(err)
	}
}
