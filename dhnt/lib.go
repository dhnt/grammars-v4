package main

import (
	"encoding/hex"
	"os"
	"bufio"
	"io/ioutil"
	"fmt"
)

//supported builtin types
//string
const builtin = `
import (
	"fmt"
)

type integer = int64
type float = float64
type object = map[interface{}]interface{}
type array = []interface{}
type boolean = bool //true false
type channel = chan interface{}
type relation = func (... interface{}) []interface{}
type null interface{}

func sprint(format string, a ...interface{}) string {
		return fmt.Sprintf(format, a)
}

func print(format string, a ...interface{}) {
		fmt.Printf(format, a)
}
`

func genLib(name string, pkg string) {
	s := fmt.Sprintf("package %s\n %s", pkg, builtin)
	d := []byte(s)
	err := ioutil.WriteFile(name, d, 0644)
	if err != nil {
		panic(err)
	}
}

// tool

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

type TargetFile struct {
	name string
	pkg  string
	imp  []string
	f    *os.File
	w    *bufio.Writer
}

func NewTargetFile(name string, pkg string, imp []string) *TargetFile {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	w.WriteString(fmt.Sprintf("package %s\n", pkg))
	w.WriteString("import (\n")
	for _, v := range imp {
		w.WriteString(v)
	}
	w.WriteString(")\n")

	return &TargetFile{name: name, pkg: pkg, imp: imp, f: f, w: w}
}

func (r *TargetFile) write(s string) {
	r.w.WriteString(s)
}

func (r *TargetFile) close() {
	r.w.Flush()
	r.f.Sync()
	r.f.Close()
}