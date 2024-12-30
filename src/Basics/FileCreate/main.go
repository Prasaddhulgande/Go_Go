package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//readeFile()
	ReadShort()
}
func createFile() {
	// Create file 
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error While creating file: ", err)
		return
	}
	defer file.Close()

	content := "Hello world By Prasad1234\n"
	//contents := "Hello world By Prasad1234\n"
	byt, errors := io.WriteString(file, content)
	fmt.Println(byt)
	if errors != nil {
		fmt.Println("Error while writeing a file")
		return
	}
	fmt.Println("Successfully created")

}

func readeFile() {

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while opening a file ", err)
		return
	}
	defer file.Close()

	// Create a buffe rto reade the file
	buffer := make([]byte, 1024)

	// Read the file content into the buffer

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while Reading the file", err)
			return
		}

		// Process the read content
		fmt.Println(string(buffer[:n]))
	}
}

func ReadShort() {
// If file is big data then this func not using cause of memory allocation some time failling use upper func.
	contetn, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while Reading a file")
		return
	}

	fmt.Println(string(contetn))
}
