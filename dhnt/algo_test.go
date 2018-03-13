package main

import (
	"testing"
)

func TestAlgo1_1(t *testing.T) {
	_, err := ParseFile("algo/1/1.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestAlgo1_2(t *testing.T) {
	_, err := ParseFile("algo/1/2.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}
