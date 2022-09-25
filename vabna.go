package main

import (
	"fmt"
	"os"
	"os/user"
	"vabna/evaluator"
	"vabna/lexer"
	"vabna/object"
	"vabna/parser"
	"vabna/repl"

	log "github.com/sirupsen/logrus"
)

func main() {

	examplecode := `
        let newAdder = fn(x) { fn(y) {x+y} };
        let addTwo = newAdder(2);
        addTwo

    `

	l := lexer.NewLexer(examplecode)
	p := parser.NewParser(&l)
	env := object.NewEnv()
	e := evaluator.Eval(p.ParseProg(), env)
	fmt.Println(e)

	//fmt.Printf("AST:\n%v\n", p.ParseProg().ToString())

	if len(p.GetErrors()) > 0 {
		var errs string

		for _, err := range p.GetErrors() {
			errs += fmt.Sprintf("%s\n", err)
		}

		log.Warnln(errs)
	}

	startRepl := true

	if startRepl {
		user, err := user.Current()

		if err != nil {
			panic(err)
		}

		fmt.Printf("Hey, %s\n", user.Username)

		repl.Repl(os.Stdin, os.Stdout)
	}

}
