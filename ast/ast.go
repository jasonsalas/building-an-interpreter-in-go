package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
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

// Program satisfies the ast.Node interface
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// LetStatement satisfies the ast.Statement & ast.Node interfaces
func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

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

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// Identifier satisfies the ast.Expression & ast.Node interfaces
func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

// Statement satisfies the ast.Node & ast.Statement interfaces
func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

// ExpressionStatement satisfies the ast.Statement & ast.Node interfaces
func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// IntegerLiteral satisfies the ast.Node & ast.Expression interfaces
func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token // the prefix token: !, -
	Operator string
	Right    Expression
}

// PrefixExpression satisfies the ast.Node & ast.Expression interfaces
func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token // the operator token: +, -, *, /, >, <, ==, !=
	Left     Expression
	Right    Expression
	Operator string
}

// InfixExpression satisfies the ast.Node & ast.Expression interfaces
func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

// Boolean satisfies the ast.Node & ast.Expression interfaces
func (b *Boolean) expressionNode() {}
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}
func (b *Boolean) String() string {
	return b.Token.Literal
}

type IfExpression struct {
	Token       token.Token // the 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

// IfExpression satisfies the ast.Node & ast.Expression interfaces
func (ie *IfExpression) expressionNode() {}
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

type BlockStatement struct {
	Token      token.Token // the '{' token
	Statements []Statement
}

// BlockStatement satisfies the ast.Node & ast.Expression interfaces
func (bs *BlockStatement) expressionNode() {}
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token // the 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

// FunctionLiteral satisfies the ast.Node & ast.Expression interfaces
func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token // the '(' token
	Function  Expression  // identifier or function literal
	Arguments []Expression
}

// CallExpression satisfies the ast.Node & ast.ExpressionNode interfaces
func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type StringLiteral struct {
	Token token.Token
	Value string
}

// StringLiteral satisfies the ast.Node & ast.Expression interfaces
func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}
func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

// ArrayLiteral satisfies the ast.Node & ast.Expression interfaces
func (al *ArrayLiteral) expressionNode() {}
func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type IndexExpression struct {
	Token token.Token // the '[' token
	Left  Expression
	Index Expression
}

// IndexExpression satisfies the ast.Node & ast.Expression interfaces
func (ie *IndexExpression) expressionNode() {}
func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

type HashLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

// HashLiteral satisfes the ast.Node & ast.Expression interfaces
func (hl *HashLiteral) expressionNode() {}
func (hl *HashLiteral) TokenLiteral() string {
	return hl.Token.Literal
}
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
