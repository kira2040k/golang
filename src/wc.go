package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName, err := getFileName()
	count_words(os.Args[1])
	if err != nil {

		os.Exit(1)
	}

	lines, err := countLinesInFile(fileName)
	if err != nil {
		os.Exit(2)
	}

	fmt.Printf("lines: %d ", lines+1)
}

func getFileName() (string, error) {
	if len(os.Args) == 1 {
		return "", errors.New(
			fmt.Sprintf("No file arg provided\nUsage: %s "+
				"[file]\n", os.Args[0]))
	}

	return os.Args[1], nil
}

func countLinesInFile(fileName string) (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return 0, err
	}

	buf := make([]byte, 1024)
	lines := 0

	for {
		readBytes, err := file.Read(buf)

		if err != nil {
			if readBytes == 0 && err == io.EOF {
				err = nil
			}
			return lines, err
		}

		lines += bytes.Count(buf[:readBytes], []byte{'\n'})
	}

	return lines, nil
}

func count_words(file_name string) {
	// initiate file-handle to read from
	fileHandle, err := os.Open(file_name)

	if err != nil {
		panic(err)
	}

	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Split(bufio.ScanWords)

	count := 0

	// for looping through results
	for fileScanner.Scan() {

		count++
	}

	// check if there was an error while reading words from file
	if err := fileScanner.Err(); err != nil {
		panic(err)
	}

	// print total word count
	fmt.Printf("words: %d \n", count)
}
