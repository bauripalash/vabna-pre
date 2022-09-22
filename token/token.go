package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

const (
    
    //Symbols
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"
    PLUS = "+"
    IDENT = "IDENT"
    INT = "INT"
    //LET = "LET"
    EQ = "="
    EQEQ = "=="
    NOT_EQ = "!="
    MUL = "*"
    DIV = "/"
    MINUS = "-"
    EXC = "!"
    
    LT = "<"
    GT = ">"
    SEMICOLON = ";"
    COMMA = ","

    //Keywords

    FUNC = "FUNCTION"
    LET = "LET"
    TRUE = "TRUE"
    FALSE = "FALSE"
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"

)

var Keywords = map[string]TokenType{
    
    "fn" : FUNC,
    "let" : LET,
    "true" : TRUE,
    "false" : FALSE,
    "if" : IF,
    "else" : ELSE,
    "return" : RETURN,
}

func LookupIdent(ident string) TokenType{
    if tok, ok := Keywords[ident]; ok{
        return tok
    }

    return IDENT
}
