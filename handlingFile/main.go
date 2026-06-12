package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Handling files in golang!")

	content := "This is a sample text to write in the file."

	file, err := os.Create("./sample.txt")

	checkNilError(err)
	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("/Users/shreyasg/Desktop/GoProgramming/handlingFile/sample.txt")
}

func readFile(filename string) {
	data, err := os.ReadFile(filename) //While reading earlier we were using ioutil but after Go 1.16, os.ReadFile is the preferred way

	checkNilError(err)

	fmt.Println("text data inside the file is: ", string(data))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
