package main

import (
	"log"
	"os"
)

func main() {
	oldLocation := os.Args[1]
	newLocation := os.Args[2]
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}
