package lexer

import (
	"github.com/ulricksennick/monkey/token"
)

// Lexer supporting ASCII characters
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	// Create a new lexer
	l := &Lexer{input: input}
	// Read the first character in the lexer's input string
	l.readChar()
	return l
}

// Generate a token from the lexer's next char, advance the lexer's position
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
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

	// Initialize token based on l.ch
	l.readChar()
	return tok
}

// Create a new token with the given token type and literal value
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Get the next character in the lexer's input, and advance the lexer's
// position in the input.
func (l *Lexer) readChar() {
	// Check if the end of the input has been reached
	if l.readPosition >= len(l.input) {
		// Set current char to 'NUL'
		l.ch = 0
	} else {
		// Set current char to next char in input
		l.ch = l.input[l.readPosition]
	}
	// Advance both position pointers
	l.position = l.readPosition
	l.readPosition++
}
