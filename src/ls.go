package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	path := flag.String("p", "./", "path for file ")
	check_all_files := flag.Bool("e", false, "check all files on directory")
	size_check := flag.Bool("l", false, "list with size")
	size := flag.String("s", "KB", "size format (kb,mb,gb)")
	flag.Parse()
	if *check_all_files {
		file_all(*path, *size)
	} else {
		defult(*path, *size, *size_check)
	}
}
func file_size(file string, size string) string {
	info, err := os.Stat(file)
	size = strings.ToUpper(size)
	if err != nil {
		return ""
	}
	if size == "KB" {

		return strconv.FormatInt(info.Size()/1000, 10) + "KB"
	}
	if size == "MB" {
		return strconv.FormatInt((info.Size()/1024)/1024, 10) + "MB"
	}
	if size == "GB" {
		return strconv.FormatInt(((info.Size()/1024)/1024)/1024, 10) + "GB"
	}
	return ""
}

func defult(path string, size string, size_check bool) {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	B := color.New(color.FgBlue)
	G := color.New(color.FgGreen, color.Bold)
	C := color.New(color.FgCyan, color.Bold)
	if size_check {

		for _, f := range files {
			info, _ := os.Stat(f.Name())

			if info.IsDir() {
				B.Print(f.Name())
				C.Print(DirSize(f.Name(), size), "\n")
			} else {
				G.Print(f.Name())
				C.Print(" " + file_size(f.Name(), size) + "\n")
			}
		}
	} else {

		for _, f := range files {
			info, _ := os.Stat(f.Name())

			if info.IsDir() {
				B.Print(f.Name() + "\n")
			} else {
				G.Print(f.Name() + "\n")

			}

		}
	}
}

func DirSize(path string, size_input string) string {
	var size int64
	size_input = strings.ToUpper(size_input)
	_ = filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	if size_input == "KB" {

		return strconv.FormatInt(size/1000, 10) + "KB"
	}
	if size_input == "MB" {
		return strconv.FormatInt((size/1024)/1024, 10) + "MB"
	}
	if size_input == "GB" {
		return strconv.FormatInt(((size/1024)/1024)/1024, 10) + "GB"
	}

	return "0"
}

func file_all(path_set string, size string) {
	var files []string
	B := color.New(color.FgBlue)
	G := color.New(color.FgGreen, color.Bold)
	C := color.New(color.FgCyan, color.Bold)
	root := path_set
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		info, _ := os.Stat(file)
		if info.IsDir() {
			B.Print(file)
			C.Print(" " + file_size(file, size) + "\n")
		} else {
			G.Print(file)
			C.Print(" " + file_size(file, size) + "\n")
		}

		//fmt.Println(file)
	}
}
