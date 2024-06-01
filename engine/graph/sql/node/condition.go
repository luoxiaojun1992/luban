package node

import (
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type Condition struct {
	Common
	Context

	LeftOperand   *commonElementsVariable.Value
	Operator      string
	RightOperands []*commonElementsVariable.Value
}
