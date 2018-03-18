package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	source := "object.jsn"

	//
	cwd, _ := os.Getwd()
	file := path.Join(cwd, "/examples/", source)

	fmt.Println("Input: ", file)

	//err := Compile(file)
	_, err := ParseFile("func.jsn")

	fmt.Printf("\nErrors: %v\n", err)

	if err != nil {
		os.Exit(1)
	}
}