package main

import (
	"fmt"
	"parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path"
	"time"
)

//
type myListener struct {
	*parser.BaseDHNTListener

	errors []string
}

func (r *myListener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("Terminal: %v\n", node)

}

func (r *myListener) VisitErrorNode(node antlr.ErrorNode) {
	fmt.Printf("ErrorNode: %v\n", node)
	r.errors = append(r.errors, fmt.Sprintf("%v\n", node))
}

type myErrorListener struct {
	*antlr.DiagnosticErrorListener
}

func (d *myErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {

}

func (d *myErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {

}

func NewErrorListener() *myErrorListener {
	return &myErrorListener{DiagnosticErrorListener: antlr.NewDiagnosticErrorListener(true)}
}

func Parse(file string) (elapsed time.Duration, err error) {
	fs, err := antlr.NewFileStream(file)
	if err != nil {
		fmt.Print(err)
		//os.Exit(1)
		return
	}

	start := time.Now()

	// Lexer
	lexer := parser.NewDHNTLexer(fs)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Parser
	p := parser.NewDHNTParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(NewErrorListener())

	// Walk the tree
	tree := p.Script()

	l := &myListener{}

	antlr.ParseTreeWalkerDefault.Walk(l, tree)

	end := time.Now()
	elapsed = end.Sub(start)

	if len(l.errors) > 0 {
		err = fmt.Errorf("Errors: %v", l.errors)
	}
	return
}

func ParseFile(ex string) (elapsed time.Duration, err error) {
	cwd, _ := os.Getwd()
	file := path.Join(cwd, "/examples/", ex)

	fmt.Println("Input: ", file)

	elapsed, err = Parse(file)

	fmt.Printf("\nErrors: %v\n", err)
	fmt.Printf("Time taken: %v\n\n", elapsed)

	return
}
//
//func main() {
//
//	_, err := ParseFile("func.jsn")
//
//	if err != nil {
//		os.Exit(1)
//	}
//}