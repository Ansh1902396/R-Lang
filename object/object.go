package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Ansh1902396/go-interpreter/ast"
)

type ObjectType string

const (
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	FUNCTION_OBJ     = "FUNCTION"
	NULL_OBJ         = "NULL"
)

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Error struct {
	Message string
}

type Integer struct {
	Value int64
}

type ReturnValue struct {
	Value Object
}
type Boolean struct {
	Value bool
}

type Null struct{}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (n *Null) Type() ObjectType {
	return NULL_OBJ
}
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

func (e *Error) Inspect() string {
	return "FUCK !! RAHULL: " + e.Message
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (n *Null) Inspect() string {
	return "null"
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}
