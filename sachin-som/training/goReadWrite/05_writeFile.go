package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// Open the output file
	outFile, outError := os.OpenFile("out2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	checkErr(outError)

	// Close
	defer outFile.Close()

	// Create bufio writer
	bufferWriter := bufio.NewWriter(outFile)

	exampleString := "dkfjdkfjkdjfkdjf\n"

	for i := 0; i < 10; i++ {
		n, err := bufferWriter.WriteString("Line No: " + string('0'+i) + " " + exampleString)
		checkErr(err)
		fmt.Println(n)
	}
	bufferWriter.Flush() // Write to the underlying io.Writer in our case outFile
}
