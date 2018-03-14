package main

import (
		"fmt"
		"parser"
		"github.com/antlr/antlr4/runtime/Go/antlr"
		"os"
		"path"
		"time"
)

type CompilationErrorListener struct {
		*antlr.DiagnosticErrorListener
}

func (d *CompilationErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {

}

func (d *CompilationErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {

}

func Compile(file string) (elapsed time.Duration, err error) {
		fs, err := antlr.NewFileStream(file)
		if err != nil {
				fmt.Print(err)
				//os.Exit(1)
				return
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

		// Walk the tree
		tree := p.Script()

		cl := NewCompilationListener(file)

		antlr.ParseTreeWalkerDefault.Walk(cl, tree)

		end := time.Now()
		elapsed = end.Sub(start)

		if len(cl.errors) > 0 {
				err = fmt.Errorf("Errors: %v", cl.errors)
		}
		return
}

func main() {
		source := "mini.jsn"

		//
		cwd, _ := os.Getwd()
		file := path.Join(cwd, "/examples/", source)

		fmt.Println("Input: ", file)

		elapsed, err := Compile(file)

		fmt.Printf("\nErrors: %v\n", err)
		fmt.Printf("Time taken: %v\n\n", elapsed)

		if err != nil {
				os.Exit(1)
		}
}