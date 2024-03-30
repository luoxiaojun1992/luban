package stmt

import (
	"fmt"

	lubanSQL "github.com/luoxiaojun1992/luban/engine/ast/sql"
)

type DeleteStmt struct {
	Table     lubanSQL.IOperand
	Condition lubanSQL.IOperand
}

func (ds *DeleteStmt) ToSQL() string {
	sqlTpl := "DELETE FROM %s %s"
	conditionSQL := ""
	if ds.Condition != nil {
		conditionTpl := "WHERE %s"
		conditionSQL = fmt.Sprintf(conditionTpl, ds.Condition.ToRaw())
	}

	return fmt.Sprint(sqlTpl, ds.Table.ToRaw(), conditionSQL)
}

func (ds *DeleteStmt) ToRaw() string {
	return ds.ToSQL()
}
