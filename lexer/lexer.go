package lexer

import "github.com/jhwagner/monkey/token"

type Lexer struct {
	input        string
	position     int  // current char position in input
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char
}

// Creates and initializes a new Lexer with a given input
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Returns the next token in the input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isValidIdentifierChar(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// Private helper function for creating token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Private helper method to read next character
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// We have either reached end of input or not read anything yet
		// Set current char to 0 (NULL)
		l.ch = 0
	} else {
		// Read next character into current char
		l.ch = l.input[l.readPosition]
	}
	// Set current position to position of char we just read and advance the read position of next char
	l.position = l.readPosition
	l.readPosition += 1
}

// Private helper method for reading in an identifier
func (l *Lexer) readIdentifier() string {
	position := l.position
	// Keep reading until we hit a non-identifier char
	for isValidIdentifierChar(l.ch) {
		l.readChar()
	}
	// Identifier will be slice of input from the beginning position to end position
	return l.input[position:l.position]
}

// Private function to determine if given character is allowed in identifiers
func isValidIdentifierChar(ch byte) bool {
	return 'a' <= ch && ch <= 'z' ||
		'A' <= ch && ch <= 'Z' ||
		'0' <= ch && ch <= '9' ||
		ch == '_'
}

// Private function to advance read position to next non-whitespace char
func (l *Lexer) skipWhitespace() {
	// If current char is a space, tab, newline, etc. read in next char
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
