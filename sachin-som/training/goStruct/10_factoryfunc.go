/*
factory functions are constructor like
functions to instansiate new Structs.
Usually starts with New or new keywords.
*/

package main

import (
	"fmt"
	"unsafe"
)

type File struct {
	fd        int
	filenname string
}

/*
Prefer way is to use NewFile as newFile can't use used inside another package
*/
func NewFile(fd int, filename string) (file *File) {
	if fd < 0 || filename == "" {
		return nil
	}
	return &File{fd: fd, filenname: filename}
}

func main() {
	file1 := NewFile(1, "text.txt")
	size := unsafe.Sizeof(File{})
	fmt.Println(file1, size)
}
