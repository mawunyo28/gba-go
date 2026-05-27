package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	var file string

	// let user enter file path and search

	fmt.Print("Enter a file relative: ")

	fmt.Scanln(&file)

	data, err := os.ReadFile(file)

	check(err)

	fmt.Println(len(data))
}
