package main

import (
	"fmt"
	"os"
	"os/user"
	"vabna/lexer"
	"vabna/parser"
	"vabna/repl"

	log "github.com/sirupsen/logrus"
)

func main() {

	examplecode := `
        fn() { return a+b; }
    `

	l := lexer.NewLexer(examplecode)
	p := parser.NewParser(&l)
	fmt.Printf("AST:\n%v\n", p.ParseProg().ToString())

	if len(p.GetErrors()) > 0 {
		var errs string

		for _, err := range p.GetErrors() {
			errs += fmt.Sprintf("%s\n", err)
		}

		log.Warnln(errs)
	}

	startRepl := false

	if startRepl {
		user, err := user.Current()

		if err != nil {
			panic(err)
		}

		fmt.Printf("Hey, %s\n", user.Username)

		repl.Repl(os.Stdin, os.Stdout)
	}

}
