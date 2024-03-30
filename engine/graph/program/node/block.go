package node

import (
	lubanASTStmt "github.com/luoxiaojun1992/luban/engine/ast/luban/stmt"
)

type Block struct {
	CodeList []string
}

func (b *Block) ToASTBlockStmt() *lubanASTStmt.Block {
	astBodyBlock := &lubanASTStmt.Block{}

	for _, code := range b.CodeList {
		astBodyStmtExpr := &lubanASTStmt.Expr{Raw: code}
		astBodyStmtExpr.AssertStmtType()
		astBodyBlock.Stmts = append(astBodyBlock.Stmts, astBodyStmtExpr)
	}

	return astBodyBlock
}
