package ast

import (
	"bytes"

	"github.com/ulricksennick/monkey/token"
)

// Basic interface implemented by every node in the abstract syntax tree.
// The attached TokenLiteral method returns the literal value of the node's
// token, used solely for debugging and testing.
type Node interface {
	TokenLiteral() string
	String() string
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

// Get the string representation of the program node, including its statements
// (child nodes)
func (p *Program) String() string {
	var out bytes.Buffer
	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}

// Let statement AST node (let <identifier> = <expression>)
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier // identifier
	Value Expression  // expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Return string representation of the let statement node
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

// Return statement AST node (return <expression>)
type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression  // expression
}

// Return string representation of the return statement node
func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

// Expression statement AST node
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression  // expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// Return string representation of the expression node
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Identifier AST node
type Identifier struct {
	Token token.Token // token.IDENT
	Value string      // user-defined identifier
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Return string representation of the identifier node
func (i *Identifier) String() string {
	return i.Value
}

// Integer literal AST node
type IntegerLiteral struct {
	Token token.Token
	Value int64 // the integer value this node represents in source code
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }
