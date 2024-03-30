package elements

import lubanSQL "github.com/luoxiaojun1992/luban/engine/ast/sql"

type Condition struct {
	LeftOperand   lubanSQL.IOperand
	Operator      string
	RightOperands []lubanSQL.IOperand
}
