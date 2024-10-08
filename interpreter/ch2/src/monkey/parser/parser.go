package parser

import (
	"github.com/anirudhlakkaraju/go-interpreter/interpreter/ch1/src/monkey/lexer"
	"github.com/anirudhlakkaraju/go-interpreter/interpreter/ch1/src/monkey/token"
	"github.com/anirudhlakkaraju/go-interpreter/interpreter/ch2/src/monkey/ast"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read tow tokens so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
