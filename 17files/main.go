package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(fileName string) {
	dataBytes, err := os.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("Text data inside the file is\n", string(dataBytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("understanding how to work with files")

	content := "this needs to go in a file temp.txt\nis this working fine?"

	file, err := os.Create("./temp.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length if:", length)
	readFile("./temp.txt")
	file.Close()
}
