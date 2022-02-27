package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	n, err := copyFile("dst.txt", "temp.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Copy Done, %d bytes had been copied.", n)
}

func copyFile(dstFile, srcFile string) (bytesCount int64, err error) {

	// OPen the src file (io.Reader)
	src, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	// Open dst file for writting (io.Writer)
	dst, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	// Copy src --> dst
	n, copyErr := io.Copy(dst, src)
	if copyErr != nil {
		return 0, copyErr
	}
	return n, nil
}
