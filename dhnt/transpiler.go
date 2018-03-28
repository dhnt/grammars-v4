package main

import (
	"fmt"
	"parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path"
	"time"
	"strings"
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

type CompilationErrorListener struct {
	*antlr.DiagnosticErrorListener
}

func (d *CompilationErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {

}

func (d *CompilationErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {

}

//source_file target_folder
func Compile(source string, target string) error {
	fs, err := antlr.NewFileStream(source)
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
	namespace := path.Base(source)

	tree := p.Script()

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Time taken to load: %v\n\n", elapsed)

	//
	pkg := "out"
	genLib(path.Join(target, "lib.go"), pkg)

	sa := strings.Split(namespace, ".")
	outfile := NewTargetFile(path.Join(target, sa[0] + ".go"), pkg, nil)

	buildObject(outfile, namespace, tree)

	buildType(outfile, tree)

	buildFunc(outfile, tree)

	outfile.close()

	return nil
}

func buildObject(file *TargetFile, namespace string, tree antlr.Tree) {
	start := time.Now()

	l := NewObjectListener(file, namespace)

	antlr.ParseTreeWalkerDefault.Walk(l, tree)

	if len(l.errors) > 0 {
		fmt.Printf("Errors: %v", l.errors)
	}

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Time to build: %v\n\n", elapsed)
}

func buildType(file *TargetFile, tree antlr.Tree) {
	start := time.Now()

	l := NewTypeListener(file)

	antlr.ParseTreeWalkerDefault.Walk(l, tree)

	if len(l.errors) > 0 {
		fmt.Printf("Errors: %v", l.errors)
	}

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Time to build: %v\n\n", elapsed)
}

func buildFunc(file *TargetFile, tree antlr.Tree) {
	start := time.Now()

	l := NewFuncListener(file)

	antlr.ParseTreeWalkerDefault.Walk(l, tree)

	if len(l.errors) > 0 {
		fmt.Printf("Errors: %v", l.errors)
	}

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Time to build: %v\n\n", elapsed)
}