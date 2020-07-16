package token

// The type of token (identifier, keyword, etc.)
type TokenType string

// Tokens are generated from source code during lexical analysis, to be given
// to the parser. A token represents a piece of source code, and is described
// by a token type and the literal value the token represents.
type Token struct {
	Type    TokenType
	Literal string
}

// Define TokenType constants
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
