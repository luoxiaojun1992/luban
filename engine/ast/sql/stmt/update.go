package stmt

import (
	"fmt"
	"strings"

	lubanSQL "github.com/luoxiaojun1992/luban/engine/ast/sql"
	lubanSQLElements "github.com/luoxiaojun1992/luban/engine/ast/sql/elements"
)

type UpdateStmt struct {
	Table          lubanSQL.IOperand
	DataCollection []*lubanSQLElements.Data
	Condition      lubanSQL.IOperand
}

func (us *UpdateStmt) ToSQL() string {
	sqlTpl := "UPDATE %s SET %s %s"

	dataSQList := []string{}
	dataSQLTpl := "%s = %s"
	for _, data := range us.DataCollection {
		dataSQL := fmt.Sprintf(dataSQLTpl, data.Field, data.Value.ToRaw())
		dataSQList = append(dataSQList, dataSQL)
	}
	dataColSQL := strings.Join(dataSQList, ", ")

	conditionSQL := ""
	if us.Condition != nil {
		conditionTpl := "WHERE %s"
		conditionSQL = fmt.Sprintf(conditionTpl, us.Condition.ToRaw())
	}

	return fmt.Sprintf(sqlTpl, us.Table.ToRaw(), dataColSQL, conditionSQL)
}

func (us *UpdateStmt) ToRaw() string {
	return us.ToSQL()
}
