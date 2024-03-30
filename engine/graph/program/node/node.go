package node

import (
	lubanAST "github.com/luoxiaojun1992/luban/engine/ast/luban"
	commonElementsFunction "github.com/luoxiaojun1992/luban/engine/elements/function"
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type NodeType string

const (
	NodeFunc        NodeType = "func"
	NodeTwoBranch   NodeType = "two_branch"
	NodeLoop        NodeType = "loop"
	NodeMultiBranch NodeType = "multi_branch"
	NodeComponent   NodeType = "component"
)

type INode interface {
	ToASTNodeList() []lubanAST.INode
	GetName() string
	GetType() NodeType
	HasCaller() bool
	GetCaller() *commonElementsFunction.Caller
	CheckIsAsync() bool
	HasInputVars() bool
	GetInputVars() []string
	HasOutputVars() bool
	GetOutputVars() []string
	GetParams() []*commonElementsFunction.Param
	GetOutputTypes() []*commonElementsVariable.VarType
	HasPrev() bool
	GetPrev() INode
	SetPrev(prev INode)
	HasNext() bool
	GetNext() INode
	SetNext(next INode)
	HasMatchChild() bool
	GetMatchChild() INode
	SetMatchChild(ifChild INode)
	HasElseChild() bool
	GetElseChild() INode
	SetElseChild(elseChild INode)
}
