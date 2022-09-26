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

	STRING = "STRING"
	// Identifier token
	IDENT = "IDENT"

	//Left Square Bracket `[`
	LS_BRACKET = "["
	//Rigt Square Bracket `]`
	RS_BRACKET = "]"

	COLON = ":"

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

	"কাজ":    FUNC,
	"ধরি":    LET,
	"সত্য":   TRUE,
	"মিথ্যা": FALSE,
	"যদি":    IF,
	"নতুবা":  ELSE,
	"ফেরাও":  RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := Keywords[ident]; ok {
		return tok
	}

	return IDENT
}
