package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // Identifier string for - add, foo, bar, x, y, ...
	INT   = "INT"   // 12345

	// Operaters
	ASSIGN = "="
	PLUS   = "+"

	// Delimieters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// Map to store language speficic keywords
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// Returns TokenType given ident string - keyword if present in map else IDENT to indicate user-defined identifier
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
