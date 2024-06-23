package data

import (
	"github.com/luoxiaojun1992/luban/engine/graph/sql/node"
)

type NodeData struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	NodeType node.NodeType `json:"node_type"`

	Assign *Assign `json:"assign"`
	Condition *Condition `json:"condition"`
}
