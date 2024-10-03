package lexer

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (points to NEXT char after current)
	ch           byte // current char under examination
}

// Returns Lexer for input string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
