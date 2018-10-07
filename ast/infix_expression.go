package ast

import (
	"bytes"
	"fmt"

	"github.com/jcox250/monkey/token"
)

type InfixExpression struct {
	Token    token.Token // The operator token e,g, +
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExpression) expressionNode() {}

func (i *InfixExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(i.Left.String())
	out.WriteString(fmt.Sprintf(" %s ", i.Operator))
	out.WriteString(i.Right.String())
	out.WriteString(")")

	return out.String()
}
