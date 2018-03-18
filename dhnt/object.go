package main

import (
	"parser"
	"path"
	"fmt"
)

type ObjectListener struct {
	*parser.BaseDHNTListener

	errors    []string

	charname  string //raw rname
	hexname   string //hex encoded name

	namespace string
}

func (s *ObjectListener) EnterScript(ctx *parser.ScriptContext) {}

func (s *ObjectListener) ExitScript(ctx *parser.ScriptContext) {}

func (s *ObjectListener) EnterObjectMembers(ctx *parser.ObjectMembersContext) {
	fmt.Printf("type %v /* %v */struct {\n", s.hexname, s.charname)
}

func (s *ObjectListener) ExitObjectMembers(ctx *parser.ObjectMembersContext) {
	fmt.Printf("}\n")
}

func NewObjectListener(file string) *ObjectListener {
	name := path.Base(file)
	return &ObjectListener{charname: name, hexname: encodeName(name)}
}

type TypeListener struct {
	*parser.BaseDHNTListener

	errors    []string

	charname  string //raw rname
	hexname   string //hex encoded name

	namespace string
}

func NewTypeListener(file string) *TypeListener {
	name := path.Base(file)
	return &TypeListener{charname: name, hexname: encodeName(name)}
}