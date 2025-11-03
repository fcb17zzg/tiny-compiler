package ast

import (
	"bytes"
	"go-monkey-compiler/token"
)

// 每个节点都需要实现Node接口
type Node interface {
	// 返回与该节点关联的字面量(该方法仅用于调试和测试)
	TokenLiteral() string
	String() string
}

type Statement interface {
	// 实现Node接口
	Node
	// 占位方法,可以让go编译器帮忙找出误用(如Expression用Statement)
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token token.Token // token.IDENT词法单元
	Value string
}

func (i *Identifier) String() string { return i.Value }

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// let语句
type LetStatement struct {
	Token token.Token // token.Let词法单元
	Name  *Identifier // 标识符
	Value Expression  // 产生值的表达式
}

type ReturnStatement struct {
	Token       token.Token // token.RETURN词法单元
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token      token.Token // 第一个词法单元
	Expression Expression
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

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

func (rs *ReturnStatement) statementNode() {

}

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

func (es *ExpressionStatement) statementNode() {

}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
