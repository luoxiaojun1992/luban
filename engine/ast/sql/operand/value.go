package operand

import (
	"fmt"

	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type Value struct {
	commonElementsVariable.Value
}

func (v *Value) ToRaw() string {
	value := v.Value.Value
	if v.ValType.IsString() {
		value = fmt.Sprintf("\"%s\"", v.Value)
	}
	return value
}
