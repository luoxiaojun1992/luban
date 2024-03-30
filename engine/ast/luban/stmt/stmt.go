package stmt

import (
	"go/ast"

	lubanAST "github.com/luoxiaojun1992/luban/engine/ast/luban"
)

type IStmt interface {
	lubanAST.INode
	ToGoASTStmt() (ast.Stmt, error)
}
