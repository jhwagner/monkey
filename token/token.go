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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	BANG = "!"
	LT   = "<"
	GT   = ">"

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

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// Returns the type of the given identifier
func LookupIdent(ident string) TokenType {
	// Check keywords table to see if this is a reserved keyword
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	// User defined identifier
	return IDENT
}
