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

func (c *Condition) GetCondition() *Condition {
	return c
}

func (c *Condition) GetTable() string {
	return ""
}

func (c *Condition) GetField() *Field {
	return nil
}

func (c *Condition) GetAssignment() *Assign {
	return nil
}
