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

type Parser struct{
    lx *lexer.Lexer
    curTok token.Token
    peekTok token.Token

    errs []string

    prefixParseFns map[token.TokenType]prefixParseFn
    infixParseFns map[token.TokenType]infixParseFn
}

type(
    prefixParseFn func() ast.Expr
    infixParseFn func(ast.Expr) ast.Expr
)

func NewParser(l *lexer.Lexer) Parser{
     
    p:= Parser{lx: l,
    errs: []string{},
    }
    
    p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
    p.regPrefix(token.IDENT , p.parseIdent)
    p.regPrefix(token.INT , p.parseIntLit)
    p.nextToken()
    p.nextToken()
    //fmt.Println(p.curTok , p.peekTok)
    return p

}

func (p *Parser) GetErrors() []string{
    return p.errs
}

func (p *Parser) regPrefix(tokenType token.TokenType , fn prefixParseFn){
    p.prefixParseFns[tokenType] = fn
}

func (p *Parser) regInflix(tokenType token.TokenType , fn infixParseFn){
    p.infixParseFns[tokenType] = fn
}

func (p *Parser) peekErr(t token.TokenType){
    msg := fmt.Sprintf("expected token %s, but got %s instead" , t , p.peekTok.Type)

    p.errs = append(p.errs, msg)
}

func (p *Parser) nextToken(){
    p.curTok = p.peekTok
    p.peekTok = p.lx.NextToken()
}

func (p *Parser) parseIdent() ast.Expr{
    
    return &ast.Identifier{
        Token: p.curTok,
        Value: p.curTok.Literal,
    }

}

func (p *Parser) parseIntLit() ast.Expr{
    l := &ast.IntLiteral{
        Token: p.curTok,
    }

    value, err := strconv.ParseInt(p.curTok.Literal , 0 , 64)
    
    if err != nil{
        msg := fmt.Sprintf("Failed to parse %q as integer" , p.curTok.Literal)
        p.errs = append(p.errs, msg)
        return nil
    }

    l.Value = value
    return l

}

func (p *Parser) ParseProg() *ast.Program{
    prog := &ast.Program{}
    prog.Stmts = []ast.Stmt{}
    
    for p.curTok.Type != token.EOF{
        //fmt.Println(p.curTok)
        stmt := p.parseStmt()
           
        if stmt != nil{
            prog.Stmts = append(prog.Stmts, stmt)
        }

        p.nextToken()
    }

    return prog
}

func (p *Parser) parseStmt() ast.Stmt{
    //fmt.Println(p.curTok.Type , p.peekTok)
    switch p.curTok.Type{
        case token.LET:
            return p.parseLetStmt()
        case token.RETURN:
            return p.parseReturnStmt()
        default:
            return p.parseExprStmt()

    }
}


func (p *Parser) parseReturnStmt() *ast.ReturnStmt{
    stmt := &ast.ReturnStmt{Token: p.curTok}

    p.nextToken()

    for !p.isCurToken(token.SEMICOLON){
        p.nextToken()
    }
    
    log.Info(fmt.Sprintf("RETURN STMT => %v\n" , stmt))

    return stmt
}

func (p *Parser) parseLetStmt() *ast.LetStmt{
//LET <IDENTIFIER> <EQUAL_SIGN> <EXPRESSION>
    stmt := &ast.LetStmt{Token: p.curTok}
    
    if !p.peek(token.IDENT){
        return nil
    }

    stmt.Name  = ast.Identifier{Token: p.curTok , Value: p.curTok.Literal}
    if !p.peek(token.EQ){
        return nil
    }

    for !p.isCurToken(token.SEMICOLON){
        p.nextToken()
    }
    
    
    log.Info(fmt.Sprintf("LET STMT => %v\n" , stmt))
    return stmt

}

func (p *Parser) parseExprStmt() *ast.ExprStmt{
    stmt := &ast.ExprStmt{Token: p.curTok}

    stmt.Expr = p.parseExpr(LOWEST) 

    if p.isPeekToken(token.SEMICOLON){
        p.nextToken()
    }
    
    return stmt
}

func (p *Parser) parseExpr(precedence int) ast.Expr{
    prefix := p.prefixParseFns[p.curTok.Type]

    if prefix == nil{
        return nil 
    }

    leftExp := prefix()

    return leftExp
}

// Helper functions
func (p *Parser) isCurToken(t token.TokenType) bool{
// check if current token type is `t`
    return p.curTok.Type == t
}

func (p *Parser) isPeekToken(t token.TokenType) bool{
// check if next token type is `t`
    return p.peekTok.Type == t
}

func (p *Parser) peek(t token.TokenType) bool{
// checks if next token type is `t`
// and if yes, then advance to next token
    if p.isPeekToken(t){
        p.nextToken()
        return true
    }else{
        p.peekErr(t)
        return false
    }
}
