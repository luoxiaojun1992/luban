package sql

import (
	"errors"

	lubanAST "github.com/luoxiaojun1992/luban/engine/ast/luban"
	lubanSQLStmt "github.com/luoxiaojun1992/luban/engine/ast/sql/stmt"
	"github.com/luoxiaojun1992/luban/engine/graph/sql/data"
	"github.com/luoxiaojun1992/luban/engine/graph/sql/node"
)

type GraphType string

const (
	GraphInsert GraphType = "insert"
	GraphSelect GraphType = "select"
	GraphUpdate GraphType = "update"
	GraphDelete GraphType = "delete"
)

type Graph struct {
	GraphType GraphType
	NodeList  []*data.NodeData `json:"node_list"`
	EdgeList  []*Edge          `json:"edge_list"`
}

func (g *Graph) toASTInsertStmt() *lubanSQLStmt.InsertStmt {
	//todo
	return nil
}

func (g *Graph) toASTSelectStmt() *lubanSQLStmt.SelectStmt {
	//todo
	return nil
}

func (g *Graph) toASTUpdateStmt() *lubanSQLStmt.UpdateStmt {
	//todo
	return nil
}

func (g *Graph) toASTDeleteStmt() *lubanSQLStmt.DeleteStmt {
	//todo
	return nil
}

func (g *Graph) parseNode() node.INode {
	//todo
	return nil
}

func (g *Graph) ToASTStmt() (lubanSQLStmt.IStmt, error) {
	//todo parse ast func for gallery, return sql str
	switch g.GraphType {
	case GraphInsert:
		return g.toASTInsertStmt(), nil
	case GraphSelect:
		return g.toASTSelectStmt(), nil
	case GraphUpdate:
		return g.toASTUpdateStmt(), nil
	case GraphDelete:
		return g.toASTDeleteStmt(), nil
	}

	return nil, errors.New("invalid sql graph type")
}

func (g *Graph) ToAllASTNode() ([]lubanAST.INode, error) {
	//todo
	g.parseNode()
	return nil, nil
}
