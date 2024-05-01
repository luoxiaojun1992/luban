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

func (g *Graph) toASTInsertStmt() (*lubanSQLStmt.InsertStmt, error) {
	// todo
	
	startNode := g.parseNode()

        if startNode == nil {
		return nil, errors.New("missing start node")
	}

        astInsertStmt := &lubanSQLStmt.InsertStmt{}
	
	startNode.GetType()
	for {
		if !startNode.HasNext() {
			break
		}

		startNode = startNode.GetNext()
	}
	return astInsertStmt, nil
}

func (g *Graph) toASTSelectStmt() (*lubanSQLStmt.SelectStmt, error) {
	// todo
	startNode := g.parseNode()

        if startNode == nil {
		return nil, errors.New("missing start node")
	}
	
	astSelectStmt := &lubanSQLStmt.SelectStmt{}
	startNode.GetType()
	for {
		if !startNode.HasNext() {
			break
		}

		startNode = startNode.GetNext()
	}
	return astSelectStmt, nil
}

func (g *Graph) toASTUpdateStmt() (*lubanSQLStmt.UpdateStmt, error) {
	// todo
	startNode := g.parseNode()

        if startNode == nil {
		return nil, errors.New("missing start node")
	}
	
	astUpdateStmt := &lubanSQLStmt.UpdateStmt{}
	startNode.GetType()
	for {
		if !startNode.HasNext() {
			break
		}

		startNode = startNode.GetNext()
	}
	return astUpdateStmt, nil
}

func (g *Graph) toASTDeleteStmt() (*lubanSQLStmt.DeleteStmt, error) {
	// todo
	startNode := g.parseNode()

        if startNode == nil {
		return nil, errors.New("missing start node")
	}
	
	astDeleteStmt := &lubanSQLStmt.DeleteStmt{}
	startNode.GetType()
	for {
		if !startNode.HasNext() {
			break
		}

		startNode = startNode.GetNext()
	}
	return astDeleteStmt, nil
}

func (g *Graph) parseNode() node.INode {
	// todo
	return nil
}

func (g *Graph) ToASTStmt() (lubanSQLStmt.IStmt, error) {
	// todo parse ast func for gallery, return sql str
	switch g.GraphType {
	case GraphInsert:
		return g.toASTInsertStmt()
	case GraphSelect:
		return g.toASTSelectStmt()
	case GraphUpdate:
		return g.toASTUpdateStmt()
	case GraphDelete:
		return g.toASTDeleteStmt()
	}

	return nil, errors.New("invalid sql graph type")
}

func (g *Graph) ToAllASTNode() ([]lubanAST.INode, error) {
	// todo
	return nil, nil
}
