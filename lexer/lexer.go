package lexer

import (
	"vabna/token"

)

type Lexer struct{
    
    input string 
    pos int 
    readPos int 
    ch byte

}

func (l *Lexer) AtEOF() bool{
    
    if l.pos >= len(l.input){
        return true
    }

    return false
    
}

func NewLexer(input string) Lexer{
    lexer := Lexer{input: input}
    lexer.readChar()
    return lexer
}

func (l *Lexer) readChar(){
    if l.readPos >= len(l.input){
        l.ch = 0
    }else{
        l.ch = l.input[l.readPos]
    }
    l.pos = l.readPos
    l.readPos += 1
}

func (l *Lexer) NextToken() token.Token{
    
    var tk token.Token
    l.eatWhitespace() 
    switch l.ch{
        
    case '+':
        tk = NewToken(token.PLUS , l.ch)
    case '-':
        tk = NewToken(token.MINUS , l.ch)
    case '*':
        tk = NewToken(token.MUL , l.ch)
    case '/':
        tk = NewToken(token.DIV , l.ch)
    case '=':
        if l.peekChar() == '='{
            ch := l.ch
            l.readChar()
            lit := string(ch) + string(l.ch)
            tk = token.Token{Type: token.EQEQ , Literal: lit}
        }else{
            tk = NewToken(token.EQ , l.ch)
        }
    case ';':
        tk = NewToken(token.SEMICOLON , l.ch)
    case ',':
        tk = NewToken(token.COMMA , l.ch)
    case '<':
        tk = NewToken(token.LT , l.ch)
    case '>':
        tk = NewToken(token.GT, l.ch)
    case '!':
         if l.peekChar() == '='{
            ch := l.ch
            l.readChar()
            lit := string(ch) + string(l.ch)
            tk = token.Token{Type: token.NOT_EQ , Literal: lit}
        }else{
            tk = NewToken(token.EQ , l.ch)
        }
    case 0:
        tk.Literal = ""
        tk.Type = token.EOF


    default:
        if isLetter(l.ch){
            tk.Literal = l.readIdent()
            tk.Type = token.LookupIdent(tk.Literal)
            return tk
        }else if isDigit(l.ch){
            tk.Type = token.INT
            tk.Literal = l.readNum()
            return tk 
        }else{
            tk = NewToken(token.ILLEGAL , l.ch)
        }

    }
    
    l.readChar()
    return tk

}

func (l *Lexer) eatWhitespace(){
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r'{
        l.readChar()
    }
}

func NewToken(tokType token.TokenType , ch byte) token.Token{
    

    return token.Token{
        Type: tokType,
        Literal: string(ch),
    }
    
}

func (l *Lexer) readIdent() string{
        
    pos := l.pos

    for isLetter(l.ch){
        l.readChar()
    }
    return l.input[pos:l.pos]
    
}

func (l *Lexer) readNum() string{
    pos := l.pos

    for isDigit(l.ch){
        l.readChar()
    }

    return l.input[pos:l.pos]
}

func (l *Lexer) peekChar() byte{
    
    if l.readPos >= len(l.input){
        return 0
    }else{
        return l.input[l.readPos]
    }
}



func isLetter(ch byte) bool{
    return 'a'<=ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool{
    return '0' <= ch && ch <= '9'
}
