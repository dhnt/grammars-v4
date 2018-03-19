package main

import (
	"encoding/hex"
	"fmt"
)

//supported builtin types
//string
type integer = int64
type float = float64
type object = map[interface{}]interface{}
type array = []interface{}
type boolean = bool //true false
type channel = chan interface{}
type relation = func (... interface{}) []interface{}
type null interface{}

// encodeName encodes name string in hex format and prepend double underscore
func encodeName(s string) string {
	//return "__" + s
	return "__" + hex.EncodeToString([]byte(s));
}

// decodeName decodes name encoded in hex
func decodeName(s string) string {
	ds, _ := hex.DecodeString(s[2:]);
	return string(ds)
}

func sprint(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a)
}

func print(format string, a ...interface{}) {
	fmt.Printf(format, a)
}