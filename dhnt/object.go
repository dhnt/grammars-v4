package main

import (
	"parser"
	"path"
	"fmt"
	"strconv"
	"strings"
)

type ObjectListener struct {
	*parser.BaseDHNTListener

	errors   []string

	script   string

	array    bool
	channel  bool
	relation bool

	nested  int
}

func (s *ObjectListener) build(str string) {
	if s.nested > 1 {
		return
	}
	//
	if s.relation {
		return
	}
	fmt.Print(str)
}

func (s *ObjectListener) EnterScript(ctx *parser.ScriptContext) {
	s.build(fmt.Sprintf("type %v /* %v */ {\n", encodeName(s.script), s.script))
}

func (s *ObjectListener) ExitScript(ctx *parser.ScriptContext) {
	s.build("\n}\n")
}

func (s *ObjectListener) EnterObjectMembers(ctx *parser.ObjectMembersContext) {
	s.nested ++
}

func (s *ObjectListener) ExitObjectMembers(ctx *parser.ObjectMembersContext) {
	s.nested --
}

func (s *ObjectListener) EnterPair(ctx *parser.PairContext) {}

func (s *ObjectListener) ExitPair(ctx *parser.PairContext) {}


func (s *ObjectListener) EnterName(ctx *parser.NameContext) {
	s.build(fmt.Sprintf("%v  ", encodeName(ctx.GetText())))
}

func (s *ObjectListener) EnterStringLiteral(ctx *parser.StringLiteralContext) {
	s.build(fmt.Sprintf("string\n"))
}

func (s *ObjectListener) EnterIdLiteral(ctx *parser.IdentifierLiteralContext) {
	s.build(fmt.Sprintf("id\n"))
}

func (s *ObjectListener) EnterNumberLiteral(ctx *parser.NumberLiteralContext) {
	t := ctx.GetText()

	if i := strings.IndexAny(t, ".eE"); i == -1 {
		s.build(fmt.Sprintf("integer\n"))
		return
	}

	if _, err := strconv.ParseFloat(t, 64); err == nil {
		s.build(fmt.Sprintf("float\n"))
		return
	}
	panic("Invalid number: " + t)
}

func (s *ObjectListener) EnterObjectLiteral(ctx *parser.ObjectLiteralContext) {
	s.build("\n")
}

func (s *ObjectListener) EnterRelationLiteral(ctx *parser.RelationLiteralContext) {
	s.relation = true
}

func (s *ObjectListener) ExitRelationLiteral(ctx *parser.RelationLiteralContext) {
	s.relation = false
}

//
//func (s *ObjectListener) EnterArrayLiteral(ctx *parser.ArrayLiteralContext) {
//
//	fmt.Printf("[] ")
//}
//
//func (s *ObjectListener) EnterChannelLiteral(ctx *parser.ChannelLiteralContext) {
//
//	fmt.Printf("chan ")
//	s.channel = true
//}
//
//func (s *ObjectListener) ExitChannelLiteral(ctx *parser.ChannelLiteralContext) {
//
//	s.channel = false
//}

func (s *ObjectListener) EnterBooleanLiteral(ctx *parser.BooleanLiteralContext) {
	s.build(fmt.Sprintf("boolean\n"))
}

func (s *ObjectListener) EnterNullLiteral(ctx *parser.NullLiteralContext) {
	s.build(fmt.Sprintf("null\n"))
}

func NewObjectListener(file string) *ObjectListener {
	name := path.Base(file)
	return &ObjectListener{script: name}
}

type TypeListener struct {
	*parser.BaseDHNTListener

	errors   []string

	nested   int
	typeword string
	alias    string

	relation bool
}

func NewTypeListener(file string) *TypeListener {
	return &TypeListener{}
}

func (s *TypeListener) build(str string) {
	if s.relation {
		return
	}
	fmt.Print(str)
}

func (s *TypeListener) EnterObjectMembers(ctx *parser.ObjectMembersContext) {
	s.typeword = "type"
	s.alias = "="

	if s.nested > 0 {
		fmt.Printf("struct {\n")
		s.typeword = ""
		s.alias = ""
	}

	s.nested ++
}

func (s *TypeListener) ExitObjectMembers(ctx *parser.ObjectMembersContext) {
	s.nested --

	s.typeword = "type"
	s.alias = "="

	if s.nested > 0 {
		fmt.Printf("\n}\n")
	}
}

func (s *TypeListener) EnterPair(ctx *parser.PairContext) {
	//t := ctx.GetChild(i).(type)
	switch ctx.GetChild(2).(type) {
	case *parser.NameContext:
	case *parser.ValueContext:
		switch ctx.GetChild(2).(type) {
		case *parser.NameContext:
			s.relation = true
		}
	}

	fmt.Printf("enter pair: %v \n", s.relation)
}

func (s *TypeListener) ExitPair(ctx *parser.PairContext) {
	switch ctx.GetChild(2).(type) {
	case *parser.NameContext:
	case *parser.RelationLiteralContext:
		s.relation = false
	}

	fmt.Printf("exit pair: %v \n", s.relation)
}

func (s *TypeListener) EnterName(ctx *parser.NameContext) {
	s.build(fmt.Sprintf("%v %v %v ", s.typeword, encodeName(ctx.GetText()), s.alias))
}

func (s *TypeListener) EnterStringLiteral(ctx *parser.StringLiteralContext) {
	s.build("string\n")
}

func (s *TypeListener) EnterIdLiteral(ctx *parser.IdentifierLiteralContext) {
	s.build("id\n")
}

func (s *TypeListener) EnterNumberLiteral(ctx *parser.NumberLiteralContext) {
	t := ctx.GetText()

	if i := strings.IndexAny(t, ".eE"); i == -1 {
		s.build("integer\n")
		return
	}

	if _, err := strconv.ParseFloat(t, 64); err == nil {
		s.build("float\n")
		return
	}
	panic("Invalid number: " + t)
}

func (s *TypeListener) EnterBooleanLiteral(ctx *parser.BooleanLiteralContext) {
	s.build("boolean\n")
}

func (s *TypeListener) EnterNullLiteral(ctx *parser.NullLiteralContext) {
	s.build("null\n")
}

//
//func (s *TypeListener) EnterRelationLiteral(ctx *parser.RelationLiteralContext) {
//	s.relation = true
//}
//
//func (s *TypeListener) ExitRelationLiteral(ctx *parser.RelationLiteralContext) {
//	s.relation = false
//}
