package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (

	//Symbols
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	PLUS    = "+"

	// Identifier token
	IDENT = "IDENT"

	// integer
	INT = "INT"

	//Equal = sign; for assignment
	EQ = "="

	EQEQ   = "=="
	NOT_EQ = "!="
	MUL    = "*"
	DIV    = "/"
	MINUS  = "-"

	//Bang or `!`
	EXC = "!"

	LT        = "<"
	GT        = ">"
	SEMICOLON = ";"
	COMMA     = ","
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//Keywords

	FUNC   = "FUNCTION"
	LET    = "LET"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
)

var Keywords = map[string]TokenType{

	"fn":     FUNC,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := Keywords[ident]; ok {
		return tok
	}

	return IDENT
}
