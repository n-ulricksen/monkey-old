package parser

import (
	"github.com/ulricksennick/monkey/ast"
	"github.com/ulricksennick/monkey/lexer"
	"github.com/ulricksennick/monkey/token"
)

// Parser implementing recursive-descent parsing
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token // current token under examination
	peekToken token.Token // next token
}

// Create a new parser which will use the given lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// Advance the parser's pointers to current and peek tokens
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	// Create a new program node

	// Advance tokens

	// Loop over tokens until EOF

	// - Create statement based on current token (either LET, RETURN, or IF)

	// - If the new statment is not nil, push to program.Statements

	// - Advance tokens

	// Return the program
	return nil
}
