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
	l.readChar() // Initialize ch, position and readPosition
	return l
}

// Reads next char of Lexer input string
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
