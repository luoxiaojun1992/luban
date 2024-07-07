package elements

import (
	"fmt"
	"strings"
	
	lubanSQL "github.com/luoxiaojun1992/luban/engine/ast/sql"
)

type Condition struct {
	LeftOperand   lubanSQL.IOperand
	Operator      string
	RightOperands []lubanSQL.IOperand
}

func (c *Condition) ToRaw() string {
	sqlTpl := "%s %s %s"

	rightSQLTpl := "%s"
	if len(c.RightOperands) > 1 {
		rightSQLTpl = "(" + rightSQLTpl + ")"
	}
	
	var rightOperandSQLList []string
	for _, rightOperand := range c.RightOperands {
		rightOperandSQLList = append(rightOperandSQLList, rightOperand.ToRaw())
	}
	rightSQL := fmt.Sprintf(rightSQLTpl, strings.Join(rightOperandSQLList, ", "))

	return fmt.Sprintf(sqlTpl, c.LeftOperand.ToRaw(), c.Operator, rightSQL)
}
