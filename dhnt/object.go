package main

import (
	"parser"
	"fmt"
	"strconv"
	"strings"
)

/**
	struct declaration
*/
type ObjectListener struct {
	*parser.BaseDHNTListener

	file     *TargetFile

	errors   []string

	script   string

	name     string
	tval     string
	text     string

	nested   int

	relation bool
}

func NewObjectListener(file *TargetFile, name string) *ObjectListener {

	return &ObjectListener{file: file, script: name}
}

func (r *ObjectListener) print() {
	if r.nested > 1 {
		return
	}

	if r.relation {
		return
	}
	s := fmt.Sprintf("%v  %v \t\t/* %v : %v */\n", encodeName(r.name), r.tval, r.name, r.text)

	r.file.write(s)
}

func (r *ObjectListener) EnterScript(ctx *parser.ScriptContext) {
	s := fmt.Sprintf("type %v /* %v */ struct {\n", encodeName(r.script), r.script)
	r.file.write(s)
}

func (r *ObjectListener) ExitScript(ctx *parser.ScriptContext) {
	r.file.write("\n}\n")
}

func (r *ObjectListener) EnterObjectMembers(ctx *parser.ObjectMembersContext) {
	r.nested ++
}

func (r *ObjectListener) ExitObjectMembers(ctx *parser.ObjectMembersContext) {
	r.nested --
}

func (r *ObjectListener) EnterName(ctx *parser.NameContext) {
	r.name = ctx.GetText()
}

func (r *ObjectListener) EnterStringLiteral(ctx *parser.StringLiteralContext) {
	r.text = ctx.GetText()
	r.tval = "string"
	r.print()
}

func (r *ObjectListener) EnterIdentifierLiteral(ctx *parser.IdentifierLiteralContext) {
	r.text = ctx.GetText()
	r.tval = encodeName(ctx.GetText())
	r.print()
}

func (r *ObjectListener) EnterNumberLiteral(ctx *parser.NumberLiteralContext) {
	t := ctx.GetText()
	r.text = t

	if i := strings.IndexAny(t, ".eE"); i == -1 {
		r.tval = "integer"; r.print()
		return
	}

	if _, err := strconv.ParseFloat(t, 64); err == nil {
		r.tval = "float"; r.print()
		return
	}
	panic("Invalid number: " + t)
}

func (r *ObjectListener) EnterObjectLiteral(ctx *parser.ObjectLiteralContext) {
	r.text = ctx.GetText()
	r.tval = ""
	r.print()
}

func (r *ObjectListener) EnterRelationLiteral(ctx *parser.RelationLiteralContext) {
	r.relation = true
}

func (r *ObjectListener) ExitRelationLiteral(ctx *parser.RelationLiteralContext) {
	r.relation = false
}

func (r *ObjectListener) EnterArrayLiteral(ctx *parser.ArrayLiteralContext) {

}

func (r *ObjectListener) EnterChannelLiteral(ctx *parser.ChannelLiteralContext) {

}

func (r *ObjectListener) EnterBooleanLiteral(ctx *parser.BooleanLiteralContext) {
	r.text = ctx.GetText()
	r.tval = "boolean"
	r.print()
}

func (r *ObjectListener) EnterNullLiteral(ctx *parser.NullLiteralContext) {
	r.text = ctx.GetText()
	r.tval = "null"
	r.print()
}

/**
	type declaration
 */
type TypeListener struct {
	*parser.BaseDHNTListener

	file     *TargetFile

	errors   []string

	nested   int

	name     string
	tval     string
	value    string

	relation bool
}

func NewTypeListener(file *TargetFile) *TypeListener {
	return &TypeListener{file: file}
}

func (r *TypeListener) print() {
	if r.relation {
		return
	}
	s := fmt.Sprintf("type %v = %v \t\t/* %v : %v */\n", encodeName(r.name), r.tval, r.name, r.value)
	if r.nested > 1 {
		s = fmt.Sprintf("%v %v \t\t/* %v : %v */\n", encodeName(r.name), r.tval, r.name, r.value)
	}
	r.file.write(s)
}

func (r *TypeListener) EnterObjectMembers(ctx *parser.ObjectMembersContext) {
	if r.nested > 0 {
		r.tval = "struct {"
		r.value = ctx.GetText()
		r.print()
	}

	r.nested ++
}

func (r *TypeListener) ExitObjectMembers(ctx *parser.ObjectMembersContext) {
	r.nested --

	if r.nested > 0 {
		r.file.write("\n}\n")
	}
}

func (r *TypeListener) EnterName(ctx *parser.NameContext) {
	r.name = ctx.GetText()
}

func (r *TypeListener) EnterStringLiteral(ctx *parser.StringLiteralContext) {
	r.value = ctx.GetText()
	r.tval = "string"
	r.print()
}

func (r *TypeListener) EnterIdentifierLiteral(ctx *parser.IdentifierLiteralContext) {
	r.value = ctx.GetText()
	r.tval = encodeName(ctx.GetText())
	r.print()
}

func (r *TypeListener) EnterNumberLiteral(ctx *parser.NumberLiteralContext) {
	t := ctx.GetText()
	r.value = t

	if i := strings.IndexAny(t, ".eE"); i == -1 {
		r.tval = "integer"
		r.print()
		return
	}

	if _, err := strconv.ParseFloat(t, 64); err == nil {
		r.tval = "float"
		r.print()
		return
	}
	panic("Invalid number: " + t)
}

func (r *TypeListener) EnterBooleanLiteral(ctx *parser.BooleanLiteralContext) {
	r.value = ctx.GetText()
	r.tval = "boolean"
	r.print()
}

func (r *TypeListener) EnterNullLiteral(ctx *parser.NullLiteralContext) {
	r.value = ctx.GetText()
	r.tval = "null"
	r.print()
}

func (r *TypeListener) EnterRelationLiteral(ctx *parser.RelationLiteralContext) {
	r.relation = true
}

func (r *TypeListener) ExitRelationLiteral(ctx *parser.RelationLiteralContext) {
	r.relation = false
}
