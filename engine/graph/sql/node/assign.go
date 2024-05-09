package node

import (
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type Assign struct {
	Field string
	Value *commonElementsVariable.Value
}
