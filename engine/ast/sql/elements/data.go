package elements

import lubanSQL "github.com/luoxiaojun1992/luban/engine/ast/sql"

type Data struct {
	Field string
	Value lubanSQL.IOperand
}
