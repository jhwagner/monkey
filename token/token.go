package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special Types
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers/Literals
	IDENT = "IDENT" // variable/function names
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
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
