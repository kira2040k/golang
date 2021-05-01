package main

import (
	"log"
	"os"
)

func main() {

	err := os.Mkdir(os.Args[1], 0755)
	if err != nil {
		log.Fatal(err)
	}
}
