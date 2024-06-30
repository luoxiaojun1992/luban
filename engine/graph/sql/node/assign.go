package node

import (
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type Assign struct {
	Common
	Context

	Field string
	Value *commonElementsVariable.Value
}

func (a *Assign) GetAssignment() *Assign {
	return a
}

func (a *Assign) GetTable() string {
	return ""
}

func (a *Assign) GetField() *Field {
	return nil
}

func (a *Assign) GetCondition() *Condition {
	return nil
}
