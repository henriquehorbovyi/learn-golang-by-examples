package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	readFromFile()
}

func readFromFile() {
	// .:.:.: IO Util makes it easy :.:.:.
	// fileName := "data.txt"
	// bytes, _ := ioutil.ReadFile(fileName)
	// fmt.Println(string(bytes))

	// .:.:.: Reading to a byte array :.:.:.
	file, _ := os.Open("data.txt")
	fileInfo, _ := file.Stat()

	// With the [FileInfo.Size()] I create a byte array with this file's size
	bytesData := make([]byte, fileInfo.Size())
	length, _ := file.Read(bytesData)

	fmt.Println(string(bytesData[:length]))
}

func reader() {
	reader := strings.NewReader("Hello world")
	bytes := make([]byte, 3)

	for {
		length, err := reader.Read(bytes)
		fmt.Printf("byte = %s\n", bytes[:length])
		if err == io.EOF {
			break
		}
	}
}
