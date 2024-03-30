package node

import (
	"fmt"
	"strings"

	lubanAST "github.com/luoxiaojun1992/luban/engine/ast/luban"
	lubanASTFunction "github.com/luoxiaojun1992/luban/engine/ast/luban/function"
	lubanASTStmt "github.com/luoxiaojun1992/luban/engine/ast/luban/stmt"
	commonElementsFunction "github.com/luoxiaojun1992/luban/engine/elements/function"
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type TwoBranch struct {
	Common
	Context

	Cond string
}

func (tb *TwoBranch) ToASTNodeList() []lubanAST.INode {
	astNodeList := make([]lubanAST.INode, 0)

	astMainFunc := &lubanASTFunction.Function{}
	astMainFunc.Name = tb.Name

	if tb.Context.HasCaller() {
		astMainFunc.Caller = tb.Context.GetCaller().Copy()
	}
	astMainFunc.Params = make([]*commonElementsFunction.Param, 0, len(tb.Context.GetParams()))
	for _, param := range tb.Context.GetParams() {
		astMainFunc.Params = append(astMainFunc.Params, param.Copy())
	}
	astMainFunc.OutputTypes = make([]*commonElementsVariable.VarType, 0, len(tb.Context.GetOutputTypes()))
	for _, varType := range tb.Context.GetOutputTypes() {
		astMainFunc.OutputTypes = append(astMainFunc.OutputTypes, varType.Copy())
	}

	astMainFuncBodyBlock := &lubanASTStmt.Block{}

	astTwoBranch := &lubanASTStmt.TwoBranch{}
	if len(tb.Cond) > 0 {
		astTwoBranchCondExpr := &lubanASTStmt.Expr{Raw: tb.Cond}
		astTwoBranchCondExpr.AssertStmtType()
		astTwoBranch.Cond = astTwoBranchCondExpr
	}

	if tb.Context.HasMatchChild() {
		astTwoBranchMatchBlock := &lubanASTStmt.Block{}

		startNode := tb.Context.GetMatchChild()

		// varLockMap := make(map[string]*node.Variable)
		for {
			subAstNodeList := startNode.ToASTNodeList()
			astNodeList = append(astNodeList, subAstNodeList...)
			//todo support only ast func node now
			astFuncCallExpr := &lubanASTStmt.Expr{}

			//todo input & output vars, handle concurrent lock
			//var map, if asyn func, store var in map, check lock if var exists in var map
			callParams := ""
			if startNode.HasInputVars() {
				callParams = strings.Join(startNode.GetInputVars(), ", ")
			}
			astFuncCallExpr.Raw = fmt.Sprintf("%s(%s)", startNode.GetName(), callParams)
			if startNode.HasCaller() {
				astFuncCallExpr.Raw = startNode.GetCaller().VarName + "." + astFuncCallExpr.Raw
			}
			if startNode.HasOutputVars() {
				astFuncCallExpr.Raw = strings.Join(startNode.GetOutputVars(), ", ") +
					" := " + astFuncCallExpr.Raw
			}

			if startNode.CheckIsAsync() {
				astFuncCallExpr.Raw = "go " + astFuncCallExpr.Raw
			}
			astFuncCallExpr.AssertStmtType()

			astTwoBranchMatchBlock.Stmts = append(astTwoBranchMatchBlock.Stmts, astFuncCallExpr)

			if !startNode.HasNext() {
				break
			}

			startNode = startNode.GetNext()
		}

		astTwoBranch.MatchBody = astTwoBranchMatchBlock
	}

	if tb.Context.HasElseChild() {
		astTwoBranchElseBlock := &lubanASTStmt.Block{}

		startNode := tb.Context.GetElseChild()

		// varLockMap := make(map[string]*node.Variable)
		for {
			subAstNodeList := startNode.ToASTNodeList()
			astNodeList = append(astNodeList, subAstNodeList...)
			//todo support only ast func node now
			astFuncCallExpr := &lubanASTStmt.Expr{}

			//todo input & output vars, handle concurrent lock
			//var map, if asyn func, store var in map, check lock if var exists in var map
			callParams := ""
			if startNode.HasInputVars() {
				callParams = strings.Join(startNode.GetInputVars(), ", ")
			}
			astFuncCallExpr.Raw = fmt.Sprintf("%s(%s)", startNode.GetName(), callParams)
			if startNode.HasCaller() {
				astFuncCallExpr.Raw = startNode.GetCaller().VarName + "." + astFuncCallExpr.Raw
			}
			if startNode.HasOutputVars() {
				astFuncCallExpr.Raw = strings.Join(startNode.GetOutputVars(), ", ") +
					" := " + astFuncCallExpr.Raw
			}

			if startNode.CheckIsAsync() {
				astFuncCallExpr.Raw = "go " + astFuncCallExpr.Raw
			}
			astFuncCallExpr.AssertStmtType()

			astTwoBranchElseBlock.Stmts = append(astTwoBranchElseBlock.Stmts, astFuncCallExpr)

			if !startNode.HasNext() {
				break
			}

			startNode = startNode.GetNext()
		}

		astTwoBranch.ElseBody = astTwoBranchElseBlock
	}

	astMainFuncBodyBlock.Stmts = append(astMainFuncBodyBlock.Stmts, astTwoBranch)

	astMainFunc.Body = astMainFuncBodyBlock

	astNodeList = append(astNodeList, astMainFunc)

	return astNodeList
}
