package main

import (
	"fmt"
	"os"
	"os/user"

	"vabna/repl"
)

func main(){

    user, err := user.Current()

    if err != nil{
        panic(err)
    }

    fmt.Printf("Hey, %s\n" , user.Username)

    repl.Repl(os.Stdin, os.Stdout)
    
    /*
    l := lexer.NewLexer(`let age = x + 1;

    if age > 18 return
    `)
    for  !l.AtEOF() {
        fmt.Println(l.NextToken())
    
    }
    */
    //println("Vabna Lang")

}
