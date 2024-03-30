package luban

import "go/ast"

type INode interface {
	ToGoASTNode() (ast.Node, error)
}
