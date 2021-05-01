package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := flag.String("p", ".", "path to search")
	name := flag.String("name", "", "name your file")
	flag.Parse()
	file_all(*path, *name)
}

func file_all(path_set string, name string) {
	var files []string
	root := path_set
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.Contains(file, name) {
			fmt.Println(file)
		}
		//fmt.Println(file)

	}
}
