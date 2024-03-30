package stmt

import "github.com/luoxiaojun1992/luban/engine/ast/sql"

type IStmt interface {
	sql.IOperand

	ToSQL() string
}
