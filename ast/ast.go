package ast

import "farcical/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {} // marker method - "this is a Statement node"
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {} // marker method - "this is a Statement node"
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// the "x" in something like "let x = 5;"
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {} // marker method - "this is an Expression node"
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
