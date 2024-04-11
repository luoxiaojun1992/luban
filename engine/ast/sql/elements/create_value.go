package elements

import lubanSQL "github.com/luoxiaojun1992/luban/engine/ast/sql"

type CreateValue struct {
	Fields     []string
	Collection [][]lubanSQL.IOperand
}

func (cv *CreateValue) ToRaw() string {
	// todo
	return ""
}
