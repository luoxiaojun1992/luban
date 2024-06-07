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
