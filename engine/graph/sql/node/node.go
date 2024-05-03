package node

type NodeType string

const (
	NodeField     NodeType = "field"
	NodeAssign    NodeType = "assign"
	NodeCondition NodeType = "condition"
	NodeTable     NodeType = "table"
)

type INode interface {
	GetName() string
	GetType() NodeType
	HasNext() bool
	GetNext() INode

	GetTable() string
}
