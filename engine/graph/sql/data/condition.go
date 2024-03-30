package data

import (
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type Condition struct {
	LeftOperand   *commonElementsVariable.Value
	Operator      string
	RightOperands []*commonElementsVariable.Value
}
