package program

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	lubanAST "github.com/luoxiaojun1992/luban/engine/ast/luban"
	lubanASTFunction "github.com/luoxiaojun1992/luban/engine/ast/luban/function"
	lubanASTStmt "github.com/luoxiaojun1992/luban/engine/ast/luban/stmt"
	"github.com/luoxiaojun1992/luban/engine/graph/program/data"
	"github.com/luoxiaojun1992/luban/engine/graph/program/node"
)

type Graph struct {
	NodeList []*data.NodeData `json:"node_list"`
	EdgeList []*Edge          `json:"edge_list"`

	ExtraNodeList []*data.NodeData `json:"extra_node_list"`
}

func ParseJSON(jsonData string) (*Graph, error) {
	graphData := &Graph{}
	if err := json.Unmarshal([]byte(jsonData), graphData); err != nil {
		return nil, err
	}

	return graphData, nil
}

func (g *Graph) ToAllASTNode() ([]lubanAST.INode, error) {
	allASTNodeList := make([]lubanAST.INode, 0)

	astNodeList, err := g.ToASTNodeList()
	if err != nil {
		return nil, err
	}

	allASTNodeList = append(allASTNodeList, astNodeList...)

	extraASTNodeList, err := g.ToExtraASTNodeList()
	if err != nil {
		return nil, err
	}

	allASTNodeList = append(allASTNodeList, extraASTNodeList...)
	return allASTNodeList, nil
}

func (g *Graph) ToASTNodeList() ([]lubanAST.INode, error) {
	nodeIDMap := make(map[int]node.INode, len(g.NodeList))
	for _, nodeData := range g.NodeList {
		var err error
		nodeIDMap[nodeData.ID], err = nodeData.ToINode()
		if err != nil {
			return nil, err
		}
	}

	for _, edgeData := range g.EdgeList {
		fromNode, ok := nodeIDMap[edgeData.Src]
		if !ok {
			return nil, fmt.Errorf("missing source node [%d]", edgeData.Src)
		}
		toNode, ok := nodeIDMap[edgeData.Dst]
		if !ok {
			return nil, fmt.Errorf("missing destination node [%d]", edgeData.Dst)
		}
		if fromNode.GetType() == node.NodeTwoBranch {
			if edgeData.GetBranchType() == "match" {
				fromNode.SetMatchChild(toNode)
			} else {
				fromNode.SetElseChild(toNode)
			}
		} else {
			fromNode.SetNext(toNode)
		}
		toNode.SetPrev(fromNode)
	}

	var startNode node.INode
	for _, nodeInstance := range nodeIDMap {
		if !nodeInstance.HasPrev() {
			startNode = nodeInstance
			break
		}
	}

	if startNode == nil {
		return nil, errors.New("missing start node")
	}

	astNodeList := make([]lubanAST.INode, 0, 2)

	astMainFunc := &lubanASTFunction.Function{}
	astMainFunc.Name = "main"

	astMainFuncBodyBlock := &lubanASTStmt.Block{}

	// varLockMap := make(map[string]*node.Variable)
	for {
		subAstNodeList := startNode.ToASTNodeList()
		astNodeList = append(astNodeList, subAstNodeList...)
		// todo support only ast func node now
		astFuncCallExpr := &lubanASTStmt.Expr{}

		// todo input & output vars, handle concurrent lock
		// var map, if asyn func, store var in map, check lock if var exists in var map
		callParams := ""
		if startNode.HasInputVars() {
			callParams = strings.Join(startNode.GetInputVars(), ", ")
		}
		astFuncCallExpr.Raw = fmt.Sprintf("%s(%s)", startNode.GetName(), callParams)
		if startNode.HasCaller() {
			astFuncCallExpr.Raw = startNode.GetCaller().VarName + "." + astFuncCallExpr.Raw
		}
		if startNode.HasOutputVars() {
			astFuncCallExpr.Raw = strings.Join(startNode.GetOutputVars(), ", ") +
				" := " + astFuncCallExpr.Raw
		}

		if startNode.CheckIsAsync() {
			astFuncCallExpr.Raw = "go " + astFuncCallExpr.Raw
		}
		astFuncCallExpr.AssertStmtType()
		astMainFuncBodyBlock.Stmts = append(astMainFuncBodyBlock.Stmts, astFuncCallExpr)

		if !startNode.HasNext() {
			break
		}

		startNode = startNode.GetNext()
	}

	astMainFunc.Body = astMainFuncBodyBlock

	astNodeList = append(astNodeList, astMainFunc)

	return astNodeList, nil
}

func (g *Graph) ToExtraASTNodeList() ([]lubanAST.INode, error) {
	// todo
	return nil, nil
}
