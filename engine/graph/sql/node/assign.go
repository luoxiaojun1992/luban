package node

import (
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type Assign struct {
        *Context
	
	Field string
	Value *commonElementsVariable.Value
}
