package node

import (
	lubanAST "github.com/luoxiaojun1992/luban/engine/ast/luban"
	lubanASTFunction "github.com/luoxiaojun1992/luban/engine/ast/luban/function"
	lubanASTStmt "github.com/luoxiaojun1992/luban/engine/ast/luban/stmt"
	commonElementsFunction "github.com/luoxiaojun1992/luban/engine/elements/function"
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type Loop struct {
	Common
	Context

	Cond     string
	CodeList []string
}

func (l *Loop) ToASTNodeList() []lubanAST.INode {
	astNodeList := make([]lubanAST.INode, 0)

	astMainFunc := &lubanASTFunction.Function{}
	astMainFunc.Name = l.Name

	if l.Context.HasCaller() {
		astMainFunc.Caller = l.Context.GetCaller().Copy()
	}
	astMainFunc.Params = make([]*commonElementsFunction.Param, 0, len(l.Context.GetParams()))
	for _, param := range l.Context.GetParams() {
		astMainFunc.Params = append(astMainFunc.Params, param.Copy())
	}
	astMainFunc.OutputTypes = make([]*commonElementsVariable.VarType, 0, len(l.Context.GetOutputTypes()))
	for _, varType := range l.Context.GetOutputTypes() {
		astMainFunc.OutputTypes = append(astMainFunc.OutputTypes, varType.Copy())
	}

	astMainFuncBodyBlock := &lubanASTStmt.Block{}

	astLoop := &lubanASTStmt.Loop{}

	if len(l.Cond) > 0 {
		astTwoBranchCondExpr := &lubanASTStmt.Expr{Raw: l.Cond}
		astTwoBranchCondExpr.AssertStmtType()
		astLoop.Cond = astTwoBranchCondExpr
	}

	if len(l.CodeList) > 0 {
		astLoopBodyBlock := &lubanASTStmt.Block{}

		for _, code := range l.CodeList {
			astLoopBodyStmtExpr := &lubanASTStmt.Expr{Raw: code}
			astLoopBodyStmtExpr.AssertStmtType()
			astLoopBodyBlock.Stmts = append(astLoopBodyBlock.Stmts, astLoopBodyStmtExpr)
		}

		astLoop.Body = astLoopBodyBlock
	}

	astMainFuncBodyBlock.Stmts = append(astMainFuncBodyBlock.Stmts, astLoop)

	astMainFunc.Body = astMainFuncBodyBlock

	astNodeList = append(astNodeList, astMainFunc)

	return astNodeList
}
