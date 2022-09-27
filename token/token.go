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
	HOLO      = "HOLO"
	EKTI      = "EKTI"
	TAHOLE    = "TAHOLE"

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
	"kaj":    FUNC,
	"fn":     FUNC,
	"ধরি":    LET,
	"dhori":  LET,
	"let":    LET,
	"সত্য":   TRUE,
	"sotto":  TRUE,
	"মিথ্যা": FALSE,
	"mittha": FALSE,
	"যদি":    IF,
	"jodi":   IF,
	"নাহলে":  ELSE,
	"nahole": ELSE,
	"ফেরাও":  RETURN,
	"ferau":  RETURN,
	"হল":     HOLO,
	"holo":   HOLO,
	"একটি":   EKTI,
	"ekti":   EKTI,
	"তাহলে":  TAHOLE,
	"tahole": TAHOLE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := Keywords[ident]; ok {
		return tok
	}

	return IDENT
}
