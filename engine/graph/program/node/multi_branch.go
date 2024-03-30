package node

import (
	lubanAST "github.com/luoxiaojun1992/luban/engine/ast/luban"
	lubanASTFunction "github.com/luoxiaojun1992/luban/engine/ast/luban/function"
	lubanASTStmt "github.com/luoxiaojun1992/luban/engine/ast/luban/stmt"
	commonElementsFunction "github.com/luoxiaojun1992/luban/engine/elements/function"
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type MultiBranch struct {
	Common
	Context

	Needle  string     `json:"needle"`
	Matches []*Matcher `json:"matches"`
}

func (mb *MultiBranch) ToASTNodeList() []lubanAST.INode {
	astNodeList := make([]lubanAST.INode, 0)

	astMainFunc := &lubanASTFunction.Function{}
	astMainFunc.Name = mb.Name

	if mb.Context.HasCaller() {
		astMainFunc.Caller = mb.Context.GetCaller().Copy()
	}
	astMainFunc.Params = make([]*commonElementsFunction.Param, 0, len(mb.Context.GetParams()))
	for _, param := range mb.Context.GetParams() {
		astMainFunc.Params = append(astMainFunc.Params, param.Copy())
	}
	astMainFunc.OutputTypes = make([]*commonElementsVariable.VarType, 0, len(mb.Context.GetOutputTypes()))
	for _, varType := range mb.Context.GetOutputTypes() {
		astMainFunc.OutputTypes = append(astMainFunc.OutputTypes, varType.Copy())
	}

	astMainFuncBodyBlock := &lubanASTStmt.Block{}

	astMultiBranch := &lubanASTStmt.MultiBranch{}

	if len(mb.Needle) > 0 {
		astNeedleExpr := &lubanASTStmt.Expr{Raw: mb.Needle}
		astNeedleExpr.AssertStmtType()
		astMultiBranch.Needle = astNeedleExpr
	}

	for _, matcher := range mb.Matches {
		astMultiBranch.Matches = append(astMultiBranch.Matches, matcher.ToASTMatcherStmt())
	}

	astMainFuncBodyBlock.Stmts = append(astMainFuncBodyBlock.Stmts, astMultiBranch)

	astMainFunc.Body = astMainFuncBodyBlock

	astNodeList = append(astNodeList, astMainFunc)

	return astNodeList
}
