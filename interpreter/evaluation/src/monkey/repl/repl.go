package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/anirudhlakkaraju/go-interpreter/interpreter/evaluation/src/monkey/evaluator"
	"github.com/anirudhlakkaraju/go-interpreter/interpreter/evaluation/src/monkey/object"
	"github.com/anirudhlakkaraju/go-interpreter/interpreter/lexing/src/monkey/lexer"
	"github.com/anirudhlakkaraju/go-interpreter/interpreter/parsing/src/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "exit()" {
			fmt.Println("Goodbye!")
			return
		}

		if line != "" && !strings.HasSuffix(line, ";") {
			// Add a semicolon if the statement is incomplete
			line += ";"
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
