package main

import (
	"testing"
)

func TestEncodeName(t *testing.T) {
	const expected = "__4e616d65"
	encoded := encodeName("Name")

	if encoded != expected {
		t.Errorf("%v", encoded)
	}
}

func TestEncodeId(t *testing.T) {
	const expected = "_name"
	encoded := encodeId("name")

	if encoded != expected {
		t.Errorf("%v", encoded)
	}
}

func TestDecodeName(t *testing.T) {
	const expected = "Name"
	decoded := decodeName("__4e616d65")

	if decoded != expected {
		t.Errorf("%v", decoded)
	}
}

func TestDecodeId(t *testing.T) {
	const expected = "name"
	decoded := decodeId("_name")

	if decoded != expected {
		t.Errorf("%v", decoded)
	}
}