package repl

import (
	"bufio"
	"fmt"
	"io"
	"vabna/lexer"
    "vabna/parser"
)

const PROMPT = "-> "

func Repl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		input := scanner.Text()
		rlexer := lexer.NewLexer(input)

	    p := parser.NewParser(&rlexer)

        prog := p.ParseProg()

        if len(p.GetErrors()) != 0{
            showParseErrors(out , p.GetErrors())
            continue
        }

        io.WriteString(out , prog.ToString())
        io.WriteString(out , "\n")
	}
}

func showParseErrors(out io.Writer , errs []string){
    for _,msg := range errs{
        io.WriteString(out , "\t ERR >" + msg + "\n")
    }
}
