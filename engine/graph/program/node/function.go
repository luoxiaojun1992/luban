package node

import (
	lubanAST "github.com/luoxiaojun1992/luban/engine/ast/luban"
	lubanASTFunction "github.com/luoxiaojun1992/luban/engine/ast/luban/function"
	lubanASTStmt "github.com/luoxiaojun1992/luban/engine/ast/luban/stmt"
	commonElementsFunction "github.com/luoxiaojun1992/luban/engine/elements/function"
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type Function struct {
	Common
	Context

	CodeList []string
}

func (f *Function) ToASTNodeList() []lubanAST.INode {
	astNodeList := make([]lubanAST.INode, 0)

	astMainFunc := &lubanASTFunction.Function{}
	astMainFunc.Name = f.Name

	if f.Context.HasCaller() {
		astMainFunc.Caller = f.Context.GetCaller().Copy()
	}
	astMainFunc.Params = make([]*commonElementsFunction.Param, 0, len(f.Context.GetParams()))
	for _, param := range f.Context.GetParams() {
		astMainFunc.Params = append(astMainFunc.Params, param.Copy())
	}
	astMainFunc.OutputTypes = make([]*commonElementsVariable.VarType, 0, len(f.Context.GetOutputTypes()))
	for _, varType := range f.Context.GetOutputTypes() {
		astMainFunc.OutputTypes = append(astMainFunc.OutputTypes, varType.Copy())
	}

	if len(f.CodeList) > 0 {
		astMainFuncBodyBlock := &lubanASTStmt.Block{}

		for _, code := range f.CodeList {
			astFuncStmtExpr := &lubanASTStmt.Expr{Raw: code}
			astFuncStmtExpr.AssertStmtType()
			astMainFuncBodyBlock.Stmts = append(astMainFuncBodyBlock.Stmts, astFuncStmtExpr)
		}

		astMainFunc.Body = astMainFuncBodyBlock
	}

	astNodeList = append(astNodeList, astMainFunc)

	return astNodeList
}
