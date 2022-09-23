package ast

import (
	"bytes"
	"vabna/token"
)

type Node interface{
    TokenLit() string
    ToString() string
}

type Stmt interface{
    Node
    stmtNode()
}

type Expr interface{
    Node
    exprNode()
}

type Program struct{
    Stmts []Stmt
}

func (p *Program) TokenLit() string{
    
    if len(p.Stmts) > 0{
        return p.Stmts[0].TokenLit()
    }else{
        return ""
    }
    
}

func (p *Program) ToString() string{
    
    var out bytes.Buffer

    for _,stmt := range p.Stmts{
        out.WriteString(stmt.ToString())
    }

    return out.String()

}

type LetStmt struct{
    Token token.Token
    Name Identifier
    Value Expr
}

func (lst *LetStmt) stmtNode(){}
func (lst *LetStmt) TokenLit() string{
    return lst.Token.Literal
}

func (lst *LetStmt) ToString() string{
    var out bytes.Buffer

    out.WriteString(lst.TokenLit() + " ")
    out.WriteString(lst.Name.ToString())
    out.WriteString(" = ")

    if lst.Value != nil{
        out.WriteString(lst.Value.ToString())
    }

    out.WriteString(";")

    return out.String()
}

type ReturnStmt struct{
    Token token.Token
    ReturnVal Expr
}

func (r *ReturnStmt) stmtNode(){}
func (r *ReturnStmt) TokenLit() string{
    return r.Token.Literal
}

func (r *ReturnStmt) ToString() string{
    var out bytes.Buffer

    out.WriteString(r.TokenLit() + " ")
    if r.ReturnVal != nil{
        out.WriteString(r.ReturnVal.ToString())
    }
    out.WriteString(";")
    return out.String()
}

type Identifier struct{
    Token token.Token
    Value string
}

func (id *Identifier) exprNode(){}
func (id *Identifier) TokenLit() string{
    return id.Token.Literal
}

func (id *Identifier) ToString() string{

    return id.Value
}

type IntLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntLiteral) exprNode()      {}
func (il *IntLiteral) TokenLit() string { return il.Token.Literal }
func (il *IntLiteral) ToString() string       { return il.Token.Literal }


type PrefixExpr struct {
	Token token.Token
	Op string
    Right Expr
}

func (p *PrefixExpr) exprNode()      {}
func (p *PrefixExpr) TokenLit() string { return p.Token.Literal }
func (p *PrefixExpr) ToString() string{

    var out bytes.Buffer
    out.WriteString("(")
    out.WriteString(p.Op)
    out.WriteString(p.Right.ToString())
    out.WriteString(")")

    return out.String()

}



type ExprStmt struct{
    Token token.Token
    Expr Expr
}

func (e *ExprStmt) stmtNode(){}
func (e *ExprStmt) TokenLit() string { 
    return e.Token.Literal
}

func (e *ExprStmt) ToString() string{
    if e.Expr != nil{
        return e.Expr.ToString()
    }
    return ""
}
