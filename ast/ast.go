package ast

import "github.com/ulricksennick/monkey/token"

// Basic interface implemented by every node in the abstract syntax tree.
// The attached TokenLiteral method returns the literal value of the node's
// token, used solely for debugging and testing.
type Node interface {
	TokenLiteral() string
}

// Statements are nodes representing segments of code that do not produce a
// value (e.g.: let x = 5;)
type Statement interface {
	Node
	statementNode()
}

// Expressions are nodes representing segments of code that produce a value
// (e.g.: 5 + 5, 342, etc.)
// For simplicity, user-defined identifiers will also be stored as expressions.
type Expression interface {
	Node
	expressionNode()
}

// A Program node will be the root node of every AST produced by our parser
type Program struct {
	Statements []Statement
}

// Returns the literal value of the token associated with the program
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// Let statement AST node (let <identifier> = <expression>)
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier // identifier
	Value Expression  // expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier AST
type Identifier struct {
	Token token.Token // token.IDENT
	Value string      // user-defined identifier
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
