package main

import (
	"encoding/hex"
	"fmt"
)

//supported builtin types
//type string
type integer = int64
type float = float64
type object = map[interface{}]interface{}
type array = []interface{}
type boolean = bool //true false
type channel = chan interface{}
type function = func (... interface{}) []interface{}
//null

// encodeName encodes name string in hex format and prepend double underscore
func encodeName(s string) string {
	return "__" + hex.EncodeToString([]byte(s));
}

// encodeId encodes identifier by prepending single underscore
func encodeId(s string) string {
	return "_" + s
}

// decodeName decodes name encoded in hex
func decodeName(s string) string {
	ds, _ := hex.DecodeString(s[2:]);
	return string(ds)
}

// decodeId decodes by removing leading single underscore
func decodeId(s string) string {
	return string(s[1:])
}

func sprint(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a)
}

func print(format string, a ...interface{}) {
	fmt.Printf(format, a)
}