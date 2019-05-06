package main

import (
	"./imagemagick"
	maim "./maim"
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

	_, err = maim.Select()
	if err != nil {
		log.Fatal(err)
	}

	_, err = maim.Screen()
	if err != nil {
		log.Fatal(err)
	}
}
