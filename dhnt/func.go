package main

/**
	func declaration
 */

import (
	"parser"
	"fmt"
)

type FuncListener struct {
	*parser.BaseDHNTListener

	file     *TargetFile

	errors   []string

	nested   int

	name     string
	tval     string
	value    string
}

func NewFuncListener(file *TargetFile) *FuncListener {
	return &FuncListener{file: file}
}

func (r *FuncListener) print() {
	r.file.write("")
}

func (r *FuncListener) EnterName(ctx *parser.NameContext) {
	r.name = ctx.GetText()
}

func (r *FuncListener) EnterRelationLiteral(ctx *parser.RelationLiteralContext) {
}

func (r *FuncListener) ExitRelationLiteral(ctx *parser.RelationLiteralContext) {
}

func (r *FuncListener) EnterParameters(ctx *parser.ParametersContext) {
	params := ctx.AllParam()
	for _, v := range params {
		fmt.Fprintln(v.GetChildCount(), v.Get)
		fmt.Println(v.GetText())
	}
}

func (r *FuncListener) EnterResults(ctx *parser.ResultsContext) {
	results := ctx.AllResult()
	for _, v := range results {
		fmt.Println(v.GetText())
	}
}

func (r *FuncListener) EnterBlockSequence(ctx *parser.BlockSequenceContext) {}

func (r *FuncListener) EnterBlockEmpty(ctx *parser.BlockEmptyContext) {}

func (r *FuncListener) EnterSequences(ctx *parser.SequencesContext) {}


