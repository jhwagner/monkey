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
