package data

import (
	"github.com/luoxiaojun1992/luban/engine/graph/sql/node"
)

type NodeData struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	NodeType node.NodeType `json:"node_type"`

	Table *Table `json:"table"`
	Field *Field `json:"field"`
	Assignment *Assign `json:"assignment"`
	Condition *Condition `json:"condition"`
}
