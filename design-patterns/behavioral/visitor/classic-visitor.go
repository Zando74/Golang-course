package visitor

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression3(de *DoubleExpression3)
	VisitAdditionExpression3(ae *AdditionExpression3)
}

type Expression3 interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpression3 struct {
	value float64
}

func (d *DoubleExpression3) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression3(d)
}

type AdditionExpression3 struct {
	left, right Expression3
}

func (a *AdditionExpression3) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression3(a)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func (e *ExpressionPrinter) VisitDoubleExpression3(de *DoubleExpression3) {
	e.sb.WriteString(fmt.Sprintf("%g", de.value))
}

func (e *ExpressionPrinter) VisitAdditionExpression3(ae *AdditionExpression3) {
	e.sb.WriteString("(")
	ae.left.Accept(e)
	e.sb.WriteString("+")
	ae.right.Accept(e)
	e.sb.WriteString(")")
}

func NewExpression3Printer() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

func (e *ExpressionPrinter) String() string {
	return e.sb.String()
}

func main3() {
	// 1+(2+3)
	e := &AdditionExpression3{
		&DoubleExpression3{1},
		&AdditionExpression3{
			left:  &DoubleExpression3{2},
			right: &DoubleExpression3{3},
		},
	}
	ep := NewExpression3Printer()
	ep.VisitAdditionExpression3(e)
	fmt.Println(ep.String())
}
