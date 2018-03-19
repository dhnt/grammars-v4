package main

import (
	"testing"
	"fmt"
)

func TestTranspile(t *testing.T) {
	source := "object.jsn"
	//

	fmt.Println("Input: ", source)

	err := Compile(source)
	//_, err := ParseFile(source)

	if err != nil {
		t.Errorf("%v", err)
	}
}
