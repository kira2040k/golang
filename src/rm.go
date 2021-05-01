package main

import (
	"log"
	"os"
)

func main() {
	err := os.RemoveAll(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
