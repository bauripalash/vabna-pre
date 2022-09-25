package repl

import (
	"bufio"
	"fmt"
	"io"
	"vabna/evaluator"
	"vabna/lexer"
	"vabna/object"
	"vabna/parser"
)

const PROMPT = "-> "

func Repl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnv()

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

		if len(p.GetErrors()) != 0 {
			showParseErrors(out, p.GetErrors())
			continue
		}
		evals := evaluator.Eval(prog, env)
		if evals != nil {
			io.WriteString(out, evals.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func showParseErrors(out io.Writer, errs []string) {
	for _, msg := range errs {
		io.WriteString(out, "\t ERR >"+msg+"\n")
	}
}
