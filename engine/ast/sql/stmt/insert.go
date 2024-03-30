package stmt

import (
	"fmt"

	lubanSQL "github.com/luoxiaojun1992/luban/engine/ast/sql"
)

type InsertStmt struct {
	Table lubanSQL.IOperand
	Data  lubanSQL.IOperand
}

func (is *InsertStmt) ToSQL() string {
	sqlTpl := "INSERT INTO %s %s"
	return fmt.Sprintf(sqlTpl, is.Table.ToRaw(), is.Data.ToRaw())
}

func (is *InsertStmt) ToRaw() string {
	return is.ToSQL()
}
