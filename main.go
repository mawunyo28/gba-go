package main

import (
	"fmt"
	"os"

	"github.com/mawunyo28/gba-go/cartridge"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	args := os.Args[1:]

	file := args[0]

	data, err := os.ReadFile(file)

	if err != nil {

		fmt.Errorf("Unable to read file")
	}

	cartridge, err := cartridge.New(data)

	if err != nil {
		fmt.Errorf("Unable to read cartridge %v", err.Error())
	}

	// let user enter file path and search

	 
	fmt.Println("Rom cartridge title: ", string(cartridge.Title[:]))

	check(err)

	fmt.Println(len(data))
}
