package main

import (
	"fmt"

	"io"
	"os"
)

func main() {

	filesystem := os.DirFS("Documents")

	file, err := filesystem.Open("note.txt")

	if err != nil {

		fmt.Println("file not found", err)

	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println(string(data))

}
