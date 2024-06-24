package visitor

import (
	"fmt"
	"strings"
)

type Expression2 interface {
	// nothing here!
}

type DoubleExpression2 struct {
	value float64
}

type AdditionExpression2 struct {
	left, right Expression2
}

func Print(e Expression2, sb *strings.Builder) {
	if de, ok := e.(*DoubleExpression2); ok {
		sb.WriteString(fmt.Sprintf("%g", de.value))
	} else if ae, ok := e.(*AdditionExpression2); ok {
		sb.WriteString("(")
		Print(ae.left, sb)
		sb.WriteString("+")
		Print(ae.right, sb)
		sb.WriteString(")")
	}

	// breaks OCP
	// will work incorrectly on missing case
}

func main2() {
	// 1+(2+3)
	e := &AdditionExpression2{
		&DoubleExpression2{1},
		&AdditionExpression2{
			left:  &DoubleExpression2{2},
			right: &DoubleExpression2{3},
		},
	}
	sb := strings.Builder{}
	Print(e, &sb)
	fmt.Println(sb.String())
}
