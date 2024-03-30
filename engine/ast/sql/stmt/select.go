package stmt

import (
	"fmt"
	"strings"

	lubanSQL "github.com/luoxiaojun1992/luban/engine/ast/sql"
	lubanSQLElements "github.com/luoxiaojun1992/luban/engine/ast/sql/elements"
)

type SelectStmt struct {
	Table     lubanSQL.IOperand
	Fields    []*lubanSQLElements.Field
	Condition lubanSQL.IOperand
}

func (ss *SelectStmt) ToSQL() string {
	tpl := "select %s from %s"

	fromSQL := ""
	if len(ss.Fields) > 0 {
		fromTpl := "(%s)"
		fieldSQLList := []string{}
		for _, field := range ss.Fields {
			fieldSQLList = append(fieldSQLList, field.ToRaw())
		}
		fromSQL = fmt.Sprintf(fromTpl, strings.Join(fieldSQLList, ", "))
	}

	return fmt.Sprintf(tpl, fromSQL, ss.Table.ToRaw())
}

func (ss *SelectStmt) ToRaw() string {
	return ss.ToSQL()
}
