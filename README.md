# Writing an Interpreter in Go

Welcome to my implementation of an interpreter for the [Monkey Programming Language](https://monkeylang.org/) written Go.

I have referenced the book *Writing an Interpreter in Go* by Thorsten Ball. Read the [book](https://interpreterbook.com/) - it's incredible!

## Overview

The interpreter is built in three steps:  

- **Lexing**: Converting the given input program string into tokens. 
- **Parsing**: The program is parsed into an Abstract Syntax Tree (AST) using Pratt Parsing approach.
- **Evaluation**: The program is finally executed by "Walking" the AST.

The interpreter supports `functions`, allowing users to define and invoke them with parameters, with global and local scoping. It also handles `strings`, `arrays` and `hashes` with respective built-in functions - `len`, `puts`, `first`, `last`, `rest`, `push`.

## Try it out! 

Clone the repository - 

```bash
$ git clone https://github.com/anirudhlakkaraju/go-interpreter.git
```

Within the `interpreter/evaluation/src/monkey` directory, run:

```bash
$ go run main.go
```

This will start the REPL (Read-Evaluate-Print-Loop). Check out the syntax for [Monkey Programming Language](https://monkeylang.org/)! 

Here are some examples of what the interpreter can do -

```bash
>> let sum = fn(x, y) {
...     x + y;
... };
>> let result = sum(5, 7);
>> result;
12
```

```bash
>> let arr = [1, "hello", fn(x) { x + 1; }];
>> let a = arr[0];
>> let b = arr[2];
>> let c = a + b(1);
3
```

```bash
>> let people = [{"name": "Alice", "age": 24}, {"name": "Anna", "age": 28}];
>> people[0]["name"];
Alice
>> people[1]["age"];
28
>> people[1]["age"] + people[0]["age"];
52
>> let getName = fn(person) { person["name"]; };
>> getName(people[0]);
Alice
>> getName(people[1]);
Anna
```

> [!NOTE]  
> Statements in Monkey Language end with a SEMICOLON (`;`).
> To exit the REPL, enter `exit()` or `CTRL-d` (i.e. `EOF`).

#### Lexer and Parser

To try out the Lexer and Parser, navigate to the `interpreter/lexer/src/monkey` directory for the Lexer and the `interpreter/parser/src/monkey` directory for the Parser, then run:

```bash
$ go run main.go
```
The Lexer REPL demonstrates how the input string is tokenized, while the Parser REPL illustrates the precedence order by correctly grouping expressions, such as converting:

```bash
>> let x = 5 + 10 / 2 * 3 - 4
let x = ((5 + ((10 / 2) * 3)) - 4)
```

## Contributions

Contributions are welcome! If you'd like to improve or extend the functionality of the interpreter, feel free to submit a PR.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
