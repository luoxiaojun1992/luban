package data

import (
	"github.com/luoxiaojun1992/luban/engine/graph/sql/node"
)

type NodeData struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	NodeType node.NodeType `json:"node_type"`

	Table      *Table     `json:"table"`
	Field      *Field     `json:"field"`
	Assignment *Assign    `json:"assignment"`
	Condition  *Condition `json:"condition"`
}

func (nd *NodeData) ToINode() (node.INode, error) {
	switch nd.NodeType {
	case node.NodeTable:
		return nd.ToTableNode()
	case node.NodeField:
		return nd.ToFieldNode()
	case node.NodeAssign:
		return nd.ToAssignNode()
	case node.NodeCondition:
		return nd.ToConditionNode()
	}
	return nil, nil
}

func (nd *NodeData) ToTableNode() (node.INode, error) {
	// todo
	return &node.Table{}, nil
}

func (nd *NodeData) ToFieldNode() (node.INode, error) {
	// todo
	return &node.Field{}, nil
}

func (nd *NodeData) ToAssignNode() (node.INode, error) {
	// todo
	return &node.Assign{}, nil
}

func (nd *NodeData) ToConditionNode() (node.INode, error) {
	// todo
	return &node.Condition{}, nil
}
