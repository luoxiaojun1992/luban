package node

import (
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type Field struct {
	Common
	Context

	Fields []*commonElementsVariable.Value
}

func (f *Field) GetField() *Field {
	return f
}

func (f *Field) GetTable() string {
	return ""
}

func (f *Field) GetAssignment() *Assign {
	return nil
}

func (f *Field) GetCondition() *Condition {
	return nil
}
