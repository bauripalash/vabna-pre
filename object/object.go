package object

import (
	"bytes"
	"fmt"
	"strings"
	"vabna/ast"
)

const (
	INT_OBJ        = "INTEGER"
	BOOL_OBJ       = "BOOLEAN"
	RETURN_VAL_OBJ = "RETURN_VAL"
	NULL_OBJ       = "NIL"
	ERR_OBJ        = "ERROR"
	FUNC_OBJ       = "FUNCTION"
)

type ObjType string

type Obj interface {
	Type() ObjType
	Inspect() string
}

//Integer 1,2,3,4,5.....100
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjType {
	return INT_OBJ
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

//Booleans true,false
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjType   { return BOOL_OBJ }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

//NULL_OBJ
type Null struct{}

func (n *Null) Type() ObjType   { return NULL_OBJ }
func (n *Null) Inspect() string { return "null" }

type ReturnValue struct {
	Value Obj
}

func (r *ReturnValue) Type() ObjType   { return RETURN_VAL_OBJ }
func (r *ReturnValue) Inspect() string { return r.Value.Inspect() }

type Error struct {
	Msg string
}

func (e *Error) Type() ObjType   { return ERR_OBJ }
func (e *Error) Inspect() string { return "ERR : " + e.Msg }

type Function struct {
	Params []*ast.Identifier
	Body   *ast.BlockStmt
	Env    *Env
}

func (f *Function) Type() ObjType { return FUNC_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range f.Params {
		params = append(params, p.ToString())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.ToString())
	out.WriteString("\n}")

	return out.String()
}
