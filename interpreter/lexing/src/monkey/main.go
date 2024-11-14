package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/anirudhlakkaraju/go-interpreter/interpreter/lexing/src/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nHello %s! This is the Monkey Programming Language!\n", user.Username)
	fmt.Println("Inputs need a semicolon to end; press Ctrl-C to quit. Synatx: https://monkeylang.org")
	fmt.Printf("\n")

	repl.Start(os.Stdin, os.Stdout)
}
