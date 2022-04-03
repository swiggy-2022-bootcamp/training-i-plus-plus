package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

type searchResult struct {
	line, file string
	lineNumber int
}

func searchFile(results chan searchResult, query, filename string) {

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Searching in file: ", filename, query)

	br := bufio.NewReader(file)
	lineNumber := 0
	for {
		lineNumber += 1
		line, _, err := br.ReadLine()
		if err != nil {
			//log.Fatalln(err)
			break
		}
		lineStr := string(line)
		// fmt.Println(lineStr)
		if strings.Contains(lineStr, query) {
			results <- searchResult{lineStr, filename, lineNumber}
		}
	}
}

func main() {

	const BasePath string = "./searchfiles"
	query := "ECG"

	files, err := ioutil.ReadDir(BasePath)

	if err != nil {
		log.Fatalln(err)
	}

	results := make(chan searchResult)
	var waitg sync.WaitGroup

	for _, f := range files {

		filePath := BasePath + "/" + f.Name()
		waitg.Add(1)
		go func() {
			defer waitg.Done()
			searchFile(results, query, filePath)
		}()
	}

	go func() {
		waitg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Printf("\nFound in file=%s @line=%d : %s\n\n", res.file, res.lineNumber, res.line)
	}

}
