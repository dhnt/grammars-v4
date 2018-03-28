package main

import (
	"testing"
	"fmt"
	"os"
	"path"
)

func TestTranspile(t *testing.T) {
	source := "object.jsn"
	//
	cwd, _ := os.Getwd()
	file := path.Join(cwd, "/examples/", source)
	target := path.Join(cwd, "/target/out/")

	os.MkdirAll(target, 0777)

	fmt.Println("Input: ", file)

	err := Compile(file, target)
	//_, err := ParseFile(source)

	if err != nil {
		t.Errorf("%v", err)
	}
}
