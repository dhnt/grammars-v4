package main

import (
		"testing"
		"os"
		"fmt"
		"path"
)

func TestTranspile(t *testing.T) {
		source := "object.jsn"
		//
		cwd, _ := os.Getwd()
		file := path.Join(cwd, "/examples/", source)

		fmt.Println("Input: ", file)

		_, err := Compile(file)

		if err != nil {
				t.Errorf("%v", err)
		}
}
