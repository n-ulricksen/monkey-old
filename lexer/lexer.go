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

	// Skip whitespace as it has no meaning, only used as separator
	l.skipWhitespace()

	switch l.ch {
	case '=':
		// If next character is '=', create a token of type EQ "=="
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '!':
		// If next character is '=', create a token of type NOT_EQ "!="
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
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
		if isLetter(l.ch) {
			// Read remaining characters of the identifier
			tok.Literal = l.readIdentifier()
			// Lookup the token type based on the identifier
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isNumber(l.ch) {
			// Read remaining characters of the INT
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			// Unrecognized character
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// Advance lexer position to next character
	l.readChar()
	return tok
}

// Skip (advance lexer) past all whitespaces including spaces, tabs, newlines,
// carriage returns
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Read an identifier from the lexer input, advancing the lexer position until
// a non-letter character is encountered
func (l *Lexer) readIdentifier() string {
	// Starting position of the identifier we are currently reading
	startPos := l.position
	// Advance lexer position until end of identifier
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

// Read a number from the lexer input, advancing the lexer position until a
// non-number character is encountered. For simplicity, integers are the only
// supported number data type.
func (l *Lexer) readNumber() string {
	// Starting position of the identifier we are currently reading
	startPos := l.position
	// Advance lexer position until end of identifier
	for isNumber(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

// Check if character is allowed in identifier (letter or underscore)
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Check if character is allowed in integer (numbers 0-9 only)
func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Create a new token with the given token type and literal value
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Advance the lexer position, updating the curent char under examination
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

// Get the next character in the lexer input, without advancing the lexer
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}
