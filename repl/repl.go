package repl

import (
	"bufio"
	"fmt"
	"io"
	"vabna/lexer"
	"vabna/token"
)


const PROMPT = "-> "

func Repl(in io.Reader , out io.Writer){
    scanner := bufio.NewScanner(in)

    for{
        fmt.Fprintf(out , PROMPT)
        scanned := scanner.Scan()

        if !scanned{
            return 
        }

        input := scanner.Text()
        rlexer := lexer.NewLexer(input)


        for tok := rlexer.NextToken(); tok.Type!= token.EOF; tok  = rlexer.NextToken(){
            fmt.Fprintf(out , "%+v\n" , tok)
        }
    }
}
