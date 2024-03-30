package node

import (
	lubanASTStmt "github.com/luoxiaojun1992/luban/engine/ast/luban/stmt"
)

type Matcher struct {
	Targets []string
	Body    *Block
}

func (m *Matcher) ToASTMatcherStmt() *lubanASTStmt.Matcher {
	astMatcher := &lubanASTStmt.Matcher{}

	for _, target := range m.Targets {
		astTargetExpr := &lubanASTStmt.Expr{Raw: target}
		astTargetExpr.AssertStmtType()
		astMatcher.Targets = append(astMatcher.Targets, astTargetExpr)
	}

	if m.Body != nil {
		astMatcher.Body = m.Body.ToASTBlockStmt()
	}

	return astMatcher
}
