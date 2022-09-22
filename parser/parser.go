package parser

import (
	"fmt"
	"vabna/ast"
	"vabna/lexer"
	"vabna/token"
    log "github.com/sirupsen/logrus"
)

type Parser struct{
    lx *lexer.Lexer
    curTok token.Token
    peekTok token.Token

    errs []string
}

func NewParser(l *lexer.Lexer) Parser{
     
    p:= Parser{lx: l,
    errs: []string{},
}
    p.nextToken()
    p.nextToken()
    //fmt.Println(p.curTok , p.peekTok)
    return p

}

func (p *Parser) GetErrors() []string{
    return p.errs
}

func (p *Parser) peekErr(t token.TokenType){
    msg := fmt.Sprintf("expected token %s, but got %s instead" , t , p.peekTok.Type)

    p.errs = append(p.errs, msg)
}

func (p *Parser) nextToken(){
    p.curTok = p.peekTok
    p.peekTok = p.lx.NextToken()
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
            return nil

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
