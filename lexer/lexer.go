package lexer

import (
	"vabna/token"
)

type Lexer struct {
	input   []rune
	pos     int
	readPos int
	ch      rune
}

func (l *Lexer) AtEOF() bool {

	return l.pos >= len(l.input)

}

/*
func getLen(inp string) int {

	return utf8.RuneCountInString(inp)

}
*/

func NewLexer(input string) Lexer {
	lexer := Lexer{input: []rune(input)}
	lexer.readChar()
	return lexer
}

func (l *Lexer) readChar() {
	//Advances lexer

	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	//fmt.Printf("<-> %c >> %d >>  %d >> %d\n", l.ch, len(string(l.ch)), l.pos, l.readPos)
	l.pos = l.readPos

	l.readPos += 1
}

func (l *Lexer) NextToken() token.Token {
	// Get next token

	var tk token.Token
	l.eatWhitespace()
	switch l.ch {

	case '+':
		tk = NewToken(token.PLUS, l.ch)
	case '-':
		tk = NewToken(token.MINUS, l.ch)
	case '*':
		tk = NewToken(token.MUL, l.ch)
	case '/':
		tk = NewToken(token.DIV, l.ch)
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			lit := string(ch) + string(l.ch)
			tk = token.Token{Type: token.EQEQ, Literal: lit}
		} else {
			tk = NewToken(token.EQ, l.ch)
		}
	case ';':
		tk = NewToken(token.SEMICOLON, l.ch)
	case ',':
		tk = NewToken(token.COMMA, l.ch)
	case '<':
		tk = NewToken(token.LT, l.ch)
	case '>':
		tk = NewToken(token.GT, l.ch)
	case '(':
		tk = NewToken(token.LPAREN, l.ch)
	case ')':
		tk = NewToken(token.RPAREN, l.ch)
	case '{':
		tk = NewToken(token.LBRACE, l.ch)
	case '}':
		tk = NewToken(token.RBRACE, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			lit := string(ch) + string(l.ch)
			tk = token.Token{Type: token.NOT_EQ, Literal: lit}
		} else {
			tk = NewToken(token.EXC, l.ch)
		}
	case '"':
		tk.Type = token.STRING
		tk.Literal = l.readString()
	case '[':
		tk = NewToken(token.LS_BRACKET, l.ch)
	case ']':
		tk = NewToken(token.RS_BRACKET, l.ch)
	case ':':
		tk = NewToken(token.COLON, l.ch)
	case 0:
		tk.Literal = ""
		tk.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tk.Literal = l.readIdent()
			tk.Type = token.LookupIdent(tk.Literal)
			return tk
		} else if isDigit(l.ch) {
			tk.Type = token.INT
			tk.Literal = l.readNum()
			return tk
		} else {
			tk = NewToken(token.ILLEGAL, l.ch)
		}

	}

	l.readChar()
	return tk

}

func (l *Lexer) readString() string {
	pos := l.pos + 1

	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	//fmt.Println(l.input[pos:l.pos])
	return string(l.input[pos:l.pos])
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func NewToken(tokType token.TokenType, ch rune) token.Token {

	return token.Token{
		Type:    tokType,
		Literal: string(ch),
	}

}

func (l *Lexer) readIdent() string {

	pos := l.pos

	for isLetter(l.ch) {
		l.readChar()
	}
	return string(l.input[pos:l.pos])

}

func (l *Lexer) readNum() string {
	pos := l.pos

	for isDigit(l.ch) {
		l.readChar()
	}

	return string(l.input[pos:l.pos])
}

func (l *Lexer) peekChar() rune {

	if l.readPos >= len(l.input) {
		return 0
	} else {
		return []rune(l.input)[l.readPos]
	}
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || 'ঀ' <= ch && ch <= '৾'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}
