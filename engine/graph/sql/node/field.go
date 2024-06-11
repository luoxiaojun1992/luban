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
