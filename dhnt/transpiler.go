package main

import (
	"fmt"
	"parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path"
	"time"
	"os"
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
	name      string
	value     string
	assertion string
}

func NewPair(name, value, assertion string) *Pair {
	return &Pair{name: name, value: value, assertion: assertion}
}

type CompilationErrorListener struct {
	*antlr.DiagnosticErrorListener
}

func (d *CompilationErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {

}

func (d *CompilationErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {

}

func Compile(source string) error {
	cwd, _ := os.Getwd()
	file := path.Join(cwd, "/examples/", source)

	fs, err := antlr.NewFileStream(file)
	if err != nil {
		fmt.Print(err)
		//os.Exit(1)
		return err
	}

	fmt.Println("source: ", fs.GetSourceName())

	start := time.Now()

	// Lexer
	lexer := parser.NewDHNTLexer(fs)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Parser
	p := parser.NewDHNTParser(stream)
	p.BuildParseTrees = true

	el := &CompilationErrorListener{DiagnosticErrorListener: antlr.NewDiagnosticErrorListener(true)}

	p.AddErrorListener(el)

	//
	namespace := path.Base(file)

	tree := p.Script()

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Time taken to load: %v\n\n", elapsed)

	buildObject(namespace, tree)

	buildType(namespace, tree)

	//buildObject(namespace, tree)

	return nil
}

func buildObject(namespace string, tree antlr.Tree) {
	start := time.Now()

	l := NewObjectListener(namespace)

	antlr.ParseTreeWalkerDefault.Walk(l, tree)

	if len(l.errors) > 0 {
		fmt.Printf("Errors: %v", l.errors)
	}

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Time to build: %v\n\n", elapsed)
}

func buildType(namespace string, tree antlr.Tree) {
	start := time.Now()

	l := NewTypeListener(namespace)

	antlr.ParseTreeWalkerDefault.Walk(l, tree)

	//if len(l.errors) > 0 {
	//	fmt.Printf("Errors: %v", l.errors)
	//}

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Time to build: %v\n\n", elapsed)
}