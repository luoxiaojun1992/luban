package data

import (
	commonElementsFunction "github.com/luoxiaojun1992/luban/engine/elements/function"
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
	"github.com/luoxiaojun1992/luban/engine/graph/program/node"
	"github.com/luoxiaojun1992/luban/engine/graph/program/node/component"
)

type NodeData struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	NodeType node.NodeType `json:"node_type"`
	IsExtra  bool          `json:"is_extra"`
	Context  *Context      `json:"context"`

	Component *Component `json:"component"`
	Function  *Function  `json:"function"`
	TwoBranch *TwoBranch `json:"two_branch"`
	Loop      *Loop      `json:"loop"`
}

func (nd *NodeData) ToINode() (node.INode, error) {
	switch nd.NodeType {
	case node.NodeFunc:
		return nd.ToFunctionNode()
	case node.NodeTwoBranch:
		return nd.ToTwoBranchNode()
	case node.NodeLoop:
		return nd.ToLoopNode()
	case node.NodeComponent:
		return nd.ToComponent()
	}
	return nil, nil
}

func (nd *NodeData) ToFunctionNode() (node.INode, error) {
	nodeFunc := &node.Function{
		Common: node.Common{
			ID:       nd.ID,
			Name:     nd.Name,
			NodeType: nd.NodeType,
		},
		Context: node.Context{},
	}

	//todo abstract
	if nd.Context != nil {
		if nd.Context.HasCaller() {
			nodeFunc.Context.Caller = nd.Context.Caller.Copy()
		}
		nodeFunc.Context.IsAsync = nd.Context.IsAsync
		nodeFunc.Context.InputVars = append(make([]string, 0, len(nd.Context.InputVars)), nd.Context.InputVars...)
		nodeFunc.Context.OutputVars = append(make([]string, 0, len(nd.Context.OutputVars)), nd.Context.OutputVars...)
		nodeFunc.Context.Params = make([]*commonElementsFunction.Param, 0, len(nd.Context.Params))
		for _, param := range nd.Context.Params {
			nodeFunc.Context.Params = append(nodeFunc.Context.Params, param.Copy())
		}
		nodeFunc.Context.OutputTypes = make([]*commonElementsVariable.VarType, 0, len(nd.Context.OutputTypes))
		for _, varType := range nd.Context.OutputTypes {
			nodeFunc.Context.OutputTypes = append(nodeFunc.Context.OutputTypes, varType.Copy())
		}
	}

	if nd.Function != nil {
		nodeFunc.CodeList = make([]string, 0, len(nd.Function.CodeList))
		for _, code := range nd.Function.CodeList {
			if !code.IsActive {
				continue
			}
			nodeFunc.CodeList = append(nodeFunc.CodeList, code.Content)
		}
	}

	return nodeFunc, nil
}

func (nd *NodeData) ToTwoBranchNode() (node.INode, error) {
	nodeTwoBranch := &node.TwoBranch{
		Common: node.Common{
			ID:       nd.ID,
			Name:     nd.Name,
			NodeType: nd.NodeType,
		},
		Context: node.Context{},
	}

	//todo abstract
	if nd.Context != nil {
		if nd.Context.HasCaller() {
			nodeTwoBranch.Context.Caller = nd.Context.Caller.Copy()
		}
		nodeTwoBranch.Context.IsAsync = nd.Context.IsAsync
		nodeTwoBranch.Context.InputVars = append(make([]string, 0, len(nd.Context.InputVars)), nd.Context.InputVars...)
		nodeTwoBranch.Context.OutputVars = append(make([]string, 0, len(nd.Context.OutputVars)), nd.Context.OutputVars...)
		nodeTwoBranch.Context.Params = make([]*commonElementsFunction.Param, 0, len(nd.Context.Params))
		for _, param := range nd.Context.Params {
			nodeTwoBranch.Context.Params = append(nodeTwoBranch.Context.Params, param.Copy())
		}
		nodeTwoBranch.Context.OutputTypes = make([]*commonElementsVariable.VarType, 0, len(nd.Context.OutputTypes))
		for _, varType := range nd.Context.OutputTypes {
			nodeTwoBranch.Context.OutputTypes = append(nodeTwoBranch.Context.OutputTypes, varType.Copy())
		}
	}

	if nd.TwoBranch != nil {
		nodeTwoBranch.Cond = nd.TwoBranch.Cond
	}

	return nodeTwoBranch, nil
}

func (nd *NodeData) ToLoopNode() (node.INode, error) {
	nodeLoop := &node.Loop{
		Common: node.Common{
			ID:       nd.ID,
			Name:     nd.Name,
			NodeType: nd.NodeType,
		},
		Context: node.Context{},
	}

	if nd.Loop != nil {
		nodeLoop.Cond = nd.Loop.Cond
		nodeLoop.CodeList = make([]string, 0, len(nd.Loop.CodeList))
		for _, code := range nd.Loop.CodeList {
			if !code.IsActive {
				continue
			}
			nodeLoop.CodeList = append(nodeLoop.CodeList, code.Content)
		}
	}

	//todo abstract
	if nd.Context != nil {
		if nd.Context.HasCaller() {
			nodeLoop.Context.Caller = nd.Context.Caller.Copy()
		}
		nodeLoop.Context.IsAsync = nd.Context.IsAsync
		nodeLoop.Context.InputVars = append(make([]string, 0, len(nd.Context.InputVars)), nd.Context.InputVars...)
		nodeLoop.Context.OutputVars = append(make([]string, 0, len(nd.Context.OutputVars)), nd.Context.OutputVars...)
		nodeLoop.Context.Params = make([]*commonElementsFunction.Param, 0, len(nd.Context.Params))
		for _, param := range nd.Context.Params {
			nodeLoop.Context.Params = append(nodeLoop.Context.Params, param.Copy())
		}
		nodeLoop.Context.OutputTypes = make([]*commonElementsVariable.VarType, 0, len(nd.Context.OutputTypes))
		for _, varType := range nd.Context.OutputTypes {
			nodeLoop.Context.OutputTypes = append(nodeLoop.Context.OutputTypes, varType.Copy())
		}
	}

	return nodeLoop, nil
}

func (nd *NodeData) ToComponent() (node.INode, error) {
	//todo
	if nd.Component != nil {
		switch nd.Component.ComponentType {
		case component.ComponentPrint:
			//todo
		}
	}
	return nil, nil
}
