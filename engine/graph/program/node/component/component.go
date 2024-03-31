package component

import (
	"errors"

	commonElementsFunction "github.com/luoxiaojun1992/luban/engine/elements/function"
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
	"github.com/luoxiaojun1992/luban/engine/graph/program/node"
)

type ComponentType string

const (
	ComponentPrint ComponentType = "print"
)

var plugins map[ComponentType]func(*ComponentInfo) node.INode

type ComponentInfo struct {
	TypeName    ComponentType
	Attrs       map[string]interface{}
	NodeID      int
	NodeName    string
	Caller      *commonElementsFunction.Caller
	IsAsync     bool
	InputVars   []string
	OutputVars  []string
	Params      []*commonElementsFunction.Param
	OutputTypes []*commonElementsVariable.VarType
}

func (ci *ComponentInfo) GetAttr(key string) interface{} {
	//todo
	return nil
}

func init() {
	plugins = make(map[ComponentType]func(*ComponentInfo) node.INode)

	AddPlugin(ComponentPrint, func(ci *ComponentInfo) node.INode {
		//todo build print code using attrs
		return &PrintComponent{
			Function: node.Function{
				Common: node.Common{
					ID:       ci.NodeID,
					Name:     ci.NodeName,
					NodeType: node.NodeComponent,
				},
				Context: node.Context{
					Caller:      ci.Caller,
					IsAsync:     ci.IsAsync,
					InputVars:   ci.InputVars,
					OutputVars:  ci.OutputVars,
					OutputTypes: ci.OutputTypes,
					Params:      ci.Params,
				},
			},
		}
	})
}

func AddPlugin(componentType ComponentType, plugin func(*ComponentInfo) node.INode) {
	plugins[componentType] = plugin
}

func CreateComponent(ci *ComponentInfo) (node.INode, error) {
	plugin, pluginExists := plugins[ci.TypeName]
	if !pluginExists {
		return nil, errors.New("plugin not exists")
	}

	return plugin(ci), nil
}
