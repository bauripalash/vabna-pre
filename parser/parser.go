package parser

import (
	"fmt"
	"strconv"
	"vabna/ast"
	"vabna/lexer"
	"vabna/token"

	log "github.com/sirupsen/logrus"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LTGT
	SUM
	PROD
	PREFIX
	CALL
)

var precedences = map[token.TokenType]int{

	token.EQ:     EQUALS,
	token.NOT_EQ: EQUALS,
	token.LT:     LTGT,
	token.GT:     LTGT,
	token.PLUS:   SUM,
	token.MINUS:  SUM,
	token.DIV:    PROD,
	token.MUL:    PROD,
}

type Parser struct {
	lx      *lexer.Lexer
	curTok  token.Token
	peekTok token.Token

	errs []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expr
	infixParseFn  func(ast.Expr) ast.Expr
)

func NewParser(l *lexer.Lexer) Parser {

	p := Parser{lx: l,
		errs: []string{},
	}

	//register prefix functions
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.regPrefix(token.IDENT, p.parseIdent)
	p.regPrefix(token.INT, p.parseIntegerLit)
	p.regPrefix(token.MINUS, p.parsePrefixExpr)
	p.regPrefix(token.EXC, p.parsePrefixExpr)

	//register infix functions
	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.regInfix(token.PLUS, p.parseInfixExpr)
	p.regInfix(token.MINUS, p.parseInfixExpr)
	p.regInfix(token.DIV, p.parseInfixExpr)
	p.regInfix(token.MUL, p.parseInfixExpr)
	p.regInfix(token.EQ, p.parseInfixExpr)
	p.regInfix(token.NOT_EQ, p.parseInfixExpr)
	p.regInfix(token.LT, p.parseInfixExpr)
	p.regInfix(token.GT, p.parseInfixExpr)
	p.nextToken()
	p.nextToken()
	//fmt.Println(p.curTok , p.peekTok)
	return p

}

func (p *Parser) GetErrors() []string {
	return p.errs
}

func (p *Parser) regPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) regInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) peekErr(t token.TokenType) {
	msg := fmt.Sprintf("expected token %s, but got %s instead", t, p.peekTok.Type)

	p.errs = append(p.errs, msg)
}

func (p *Parser) nextToken() {
	p.curTok = p.peekTok
	p.peekTok = p.lx.NextToken()
}

func (p *Parser) ParseProg() *ast.Program {
	prog := &ast.Program{}
	prog.Stmts = []ast.Stmt{}

	for p.curTok.Type != token.EOF {
		//fmt.Println(p.curTok)
		stmt := p.parseStmt()

		if stmt != nil {
			prog.Stmts = append(prog.Stmts, stmt)
		}

		p.nextToken()
	}

	return prog
}

func (p *Parser) parseStmt() ast.Stmt {
	//fmt.Println(p.curTok.Type , p.peekTok)
	switch p.curTok.Type {
	case token.LET:
		return p.parseLetStmt()
	case token.RETURN:
		return p.parseReturnStmt()
	default:
		return p.parseExprStmt()

	}
}

func (p *Parser) parseReturnStmt() *ast.ReturnStmt {
	stmt := &ast.ReturnStmt{Token: p.curTok}

	p.nextToken()

	for !p.isCurToken(token.SEMICOLON) {
		p.nextToken()
	}

	log.Info(fmt.Sprintf("RETURN STMT => %v\n", stmt))

	return stmt
}

func (p *Parser) parseLetStmt() *ast.LetStmt {
	//LET <IDENTIFIER> <EQUAL_SIGN> <EXPRESSION>
	stmt := &ast.LetStmt{Token: p.curTok}

	if !p.peek(token.IDENT) {
		return nil
	}

	stmt.Name = ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}
	if !p.peek(token.EQ) {
		return nil
	}

	for !p.isCurToken(token.SEMICOLON) {
		p.nextToken()
	}

	log.Info(fmt.Sprintf("LET STMT => %v\n", stmt))
	return stmt

}

func (p *Parser) parseExprStmt() *ast.ExprStmt {
	stmt := &ast.ExprStmt{Token: p.curTok}

	stmt.Expr = p.parseExpr(LOWEST)

	if p.isPeekToken(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) noPrefixFunctionErr(t token.TokenType) {
	msg := fmt.Sprintf("no prefix function for %s", t)
	p.errs = append(p.errs, msg)
}

func (p *Parser) parseExpr(prec int) ast.Expr {
	prefix := p.prefixParseFns[p.curTok.Type]
	if prefix == nil {
		p.noPrefixFunctionErr(p.curTok.Type)
		return nil
	}

	leftExpr := prefix()

	for !p.isPeekToken(token.SEMICOLON) && prec < p.peekPrec() {
		infix := p.infixParseFns[p.peekTok.Type]

		if infix == nil {
			return leftExpr
		}

		p.nextToken()

		leftExpr = infix(leftExpr)
	}

	return leftExpr

}

func (p *Parser) parseIdent() ast.Expr {
	log.Info("IDENT EXPR =>", p.curTok)
	return &ast.Identifier{
		Token: p.curTok,
		Value: p.curTok.Literal,
	}

}

func (p *Parser) parseIntegerLit() ast.Expr {
	log.Info("INT EXPR =>", p.curTok)
	lit := &ast.IntegerLit{Token: p.curTok}

	value, err := strconv.ParseInt(p.curTok.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("Could not parse %s as Interger", p.curTok)
		p.errs = append(p.errs, msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) parsePrefixExpr() ast.Expr {
	exp := &ast.PrefixExpr{
		Token: p.curTok,
		Op:    p.curTok.Literal,
	}

	p.nextToken()
	exp.Right = p.parseExpr(PREFIX)
	return exp
}

func (p *Parser) parseInfixExpr(left ast.Expr) ast.Expr {
	exp := &ast.InfixExpr{
		Token: p.curTok,
		Op:    p.curTok.Literal,
		Left:  left,
	}

	prec := p.curPrec()
	p.nextToken()
	exp.Right = p.parseExpr(prec)

	return exp
}

// Helper functions
func (p *Parser) isCurToken(t token.TokenType) bool {
	// check if current token type is `t`
	return p.curTok.Type == t
}

func (p *Parser) isPeekToken(t token.TokenType) bool {
	// check if next token type is `t`
	return p.peekTok.Type == t
}

func (p *Parser) peek(t token.TokenType) bool {
	// checks if next token type is `t`
	// and if yes, then advance to next token
	if p.isPeekToken(t) {
		p.nextToken()
		return true
	} else {
		p.peekErr(t)
		return false
	}
}

// Check precedence of Peek Token
//(Token after current token)
func (p *Parser) peekPrec() int {
	if p, ok := precedences[p.peekTok.Type]; ok {
		return p
	}

	return LOWEST
}

//Check precedence of Current Token
func (p *Parser) curPrec() int {
	if p, ok := precedences[p.curTok.Type]; ok {
		return p
	}
	return LOWEST
}
