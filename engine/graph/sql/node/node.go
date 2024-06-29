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
	HasPrev() bool
	GetPrev() INode
	SetPrev(prev INode)
	HasNext() bool
	GetNext() INode
	SetNext(next INode)

	GetTable() string
	GetAssignment() *Assign
	GetCondition() *Condition
}
