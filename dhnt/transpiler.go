package main

import (
		"fmt"
		"parser"
		"github.com/antlr/antlr4/runtime/Go/antlr"
		"path"
)

type Kind int

const (
		String Kind = iota
		Id
		Number
		Object
		Channel
		Function
		Array
		True
		False
		Null
)

type Pair struct {
		name string
		value string
		assertion string
}

func NewPair(name, value, assertion string) *Pair {
		return &Pair{name: name, value: value, assertion: assertion}
}

type CompilationListener struct {
		*parser.BaseDHNTListener

		errors []string

		name string
		hex string //hex encoded name

		//
		pairs []Pair

		nested uint8
}

func NewCompilationListener(file string) *CompilationListener {
		name := path.Base(file)
		return &CompilationListener{name: name, hex: encodeName(name)}
}

func (s *CompilationListener) VisitTerminal(node antlr.TerminalNode) {
		//fmt.Printf("Terminal: %v\n", node)
}

func (s *CompilationListener) VisitErrorNode(node antlr.ErrorNode) {
		fmt.Printf("ErrorNode: %v\n", node)
		s.errors = append(s.errors, fmt.Sprintf("%v\n", node))
}

func (s *CompilationListener) EnterScript(ctx *parser.ScriptContext) {
		fmt.Println("enter script")
}

func (s *CompilationListener) ExitScript(ctx *parser.ScriptContext) {
		fmt.Println("exit script")
}

func (s *CompilationListener) EnterObjectMembers(ctx *parser.ObjectMembersContext) {
		if s.nested == 0 {
				fmt.Printf("type %v /* %v */ struct {\n", s.hex, s.name)
		} else {
				fmt.Printf("type %v /* %v */ struct {\n", ctx.GetChildren(), ctx.GetText())
		}
		s.nested ++
}

func (s *CompilationListener) ExitObjectMembers(ctx *parser.ObjectMembersContext) {
		fmt.Printf("}\n")
		s.nested --
}

func (s *CompilationListener) EnterObjectKind(ctx *parser.ObjectKindContext) {}

func (s *CompilationListener) ExitObjectKind(ctx *parser.ObjectKindContext) {}

func (s *CompilationListener) EnterObjectZero(ctx *parser.ObjectZeroContext) {}

func (s *CompilationListener) ExitObjectZero(ctx *parser.ObjectZeroContext) {}

func (s *CompilationListener) EnterArrayElements(ctx *parser.ArrayElementsContext) {}

func (s *CompilationListener) ExitArrayElements(ctx *parser.ArrayElementsContext) {}

func (s *CompilationListener) EnterArrayKind(ctx *parser.ArrayKindContext) {}

func (s *CompilationListener) ExitArrayKind(ctx *parser.ArrayKindContext) {}

func (s *CompilationListener) EnterArrayZero(ctx *parser.ArrayZeroContext) {}

func (s *CompilationListener) ExitArrayZero(ctx *parser.ArrayZeroContext) {}

func (s *CompilationListener) EnterFunction(ctx *parser.FunctionContext) {}

func (s *CompilationListener) ExitFunction(ctx *parser.FunctionContext) {}

func (s *CompilationListener) EnterPair(ctx *parser.PairContext) {

}

func (s *CompilationListener) ExitPair(ctx *parser.PairContext) {

		fmt.Printf("exit pair: %v\n", ctx.GetText())
}

func (s *CompilationListener) EnterName(ctx *parser.NameContext) {}

func (s *CompilationListener) ExitName(ctx *parser.NameContext) {}

func (s *CompilationListener) EnterLiteralValue(ctx *parser.LiteralValueContext) {}

func (s *CompilationListener) ExitLiteralValue(ctx *parser.LiteralValueContext) {}

func (s *CompilationListener) EnterExpressionValue(ctx *parser.ExpressionValueContext) {}

func (s *CompilationListener) ExitExpressionValue(ctx *parser.ExpressionValueContext) {}

func (s *CompilationListener) EnterKind(ctx *parser.KindContext) {}

func (s *CompilationListener) ExitKind(ctx *parser.KindContext) {}

func (s *CompilationListener) EnterLiteral(ctx *parser.LiteralContext) {}

func (s *CompilationListener) ExitLiteral(ctx *parser.LiteralContext) {}
