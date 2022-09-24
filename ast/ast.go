package ast

import (
	"bytes"
	"strings"
	"vabna/token"
)

type Node interface {
	TokenLit() string
	ToString() string
}

type Stmt interface {
	Node
	stmtNode()
}

type Expr interface {
	Node
	exprNode()
}

//Program - main entry point

type Program struct {
	Stmts []Stmt
}

func (p *Program) TokenLit() string {

	if len(p.Stmts) > 0 {
		return p.Stmts[0].TokenLit()
	} else {
		return ""
	}

}

func (p *Program) ToString() string {

	var out bytes.Buffer

	for _, stmt := range p.Stmts {
		out.WriteString(stmt.ToString())
	}

	return out.String()

}

// let statement

type LetStmt struct {
	Token token.Token
	Name  Identifier
	Value Expr
}

func (lst *LetStmt) stmtNode() {}
func (lst *LetStmt) TokenLit() string {
	return lst.Token.Literal
}

func (lst *LetStmt) ToString() string {
	var out bytes.Buffer

	out.WriteString(lst.TokenLit() + " ")
	out.WriteString(lst.Name.ToString())
	out.WriteString(" = ")

	if lst.Value != nil {
		out.WriteString(lst.Value.ToString())
	}

	out.WriteString(";")

	return out.String()
}

//Return statement

type ReturnStmt struct {
	Token     token.Token
	ReturnVal Expr
}

func (r *ReturnStmt) stmtNode() {}
func (r *ReturnStmt) TokenLit() string {
	return r.Token.Literal
}

func (r *ReturnStmt) ToString() string {
	var out bytes.Buffer

	out.WriteString(r.TokenLit() + " ")
	if r.ReturnVal != nil {
		out.WriteString(r.ReturnVal.ToString())
	}
	out.WriteString(";")
	return out.String()
}

//Expression Statment
type ExprStmt struct {
	Token token.Token
	Expr  Expr
}

func (e *ExprStmt) stmtNode() {}
func (e *ExprStmt) TokenLit() string {
	return e.Token.Literal
}

func (e *ExprStmt) ToString() string {
	//fmt.Println(e.Expr.TokenLit())
	if e.Expr != nil {
		return e.Expr.ToString()
	} else {

		return ""
	}
}

// Identifier Expression

type Identifier struct {
	Token token.Token
	Value string
}

func (id *Identifier) exprNode() {}
func (id *Identifier) TokenLit() string {
	return id.Token.Literal
}

func (id *Identifier) ToString() string {

	return id.Value
}

//Integer Expression

type IntegerLit struct {
	Token token.Token
	Value int64
}

func (in *IntegerLit) exprNode() {}
func (in *IntegerLit) TokenLit() string {
	return in.Token.Literal
}

func (in *IntegerLit) ToString() string {
	return in.Token.Literal
}

// Prefix Expression

type PrefixExpr struct {
	Token token.Token
	Op    string
	Right Expr
}

func (pref *PrefixExpr) exprNode()        {}
func (pref *PrefixExpr) TokenLit() string { return pref.Token.Literal }
func (pref *PrefixExpr) ToString() string {

	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pref.Op)
	out.WriteString(pref.Right.ToString())
	out.WriteString(")")
	return out.String()

}

//Infix Expression

type InfixExpr struct {
	Token token.Token
	Left  Expr
	Op    string
	Right Expr
}

func (inf *InfixExpr) exprNode()        {}
func (inf *InfixExpr) TokenLit() string { return inf.Token.Literal }
func (inf *InfixExpr) ToString() string {

	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(inf.Left.ToString())
	out.WriteString(" " + inf.Op + " ")
	out.WriteString(inf.Right.ToString())
	out.WriteString(")")
	return out.String()

}

//Boolean Expression
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) exprNode()        {}
func (b *Boolean) TokenLit() string { return b.Token.Literal }
func (b *Boolean) ToString() string { return b.Token.Literal }

type BlockStmt struct {
	Token token.Token
	Stmts []Stmt
}

func (bs *BlockStmt) stmtNode()        {}
func (bs *BlockStmt) TokenLit() string { return bs.Token.Literal }
func (bs *BlockStmt) ToString() string {
	var out bytes.Buffer
	for _, s := range bs.Stmts {
		out.WriteString(s.ToString())
	}
	return out.String()
}

type IfExpr struct {
	Token     token.Token
	Cond      Expr
	TrueBlock *BlockStmt
	ElseBlock *BlockStmt
}

func (i *IfExpr) exprNode()        {}
func (i *IfExpr) TokenLit() string { return i.Token.Literal }
func (i *IfExpr) ToString() string {

	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(i.Cond.ToString())
	out.WriteString(" ")
	out.WriteString(i.TrueBlock.ToString())
	if i.ElseBlock != nil {

		out.WriteString("else ")
		out.WriteString(i.ElseBlock.ToString())
	}
	return out.String()

}

type FunctionLit struct {
	Token  token.Token // The 'fn' token
	Params []*Identifier
	Body   *BlockStmt
}

func (fl *FunctionLit) exprNode()        {}
func (fl *FunctionLit) TokenLit() string { return fl.Token.Literal }
func (fl *FunctionLit) ToString() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Params {
		params = append(params, p.ToString())
	}
	out.WriteString(fl.TokenLit())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.ToString())
	return out.String()
}

type CallExpr struct {
	Token token.Token // The '(' token
	Func  Expr
	// Identifier or FunctionLiteral
	Args []Expr
}

func (ce *CallExpr) exprNode()        {}
func (ce *CallExpr) TokenLit() string { return ce.Token.Literal }
func (ce *CallExpr) ToString() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Args {
		args = append(args, a.ToString())
	}
	out.WriteString(ce.Func.ToString())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}
