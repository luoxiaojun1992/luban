package operand

import (
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type Value struct {
	commonElementsVariable.Value
}

func (v *Value) ToRaw() string {
	return v.ToString()
}
