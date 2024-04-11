package graph

import (
	"encoding/json"

	lubanAST "github.com/luoxiaojun1992/luban/engine/ast/luban"
	"github.com/luoxiaojun1992/luban/engine/graph/program"
)

type GraphType string

const (
	GraphProgram GraphType = "program"
	GraphSQL     GraphType = "sql"
)

type IGraph interface {
	ToAllASTNode() ([]lubanAST.INode, error)
}

type Graph struct {
	GraphType GraphType `json:"graph_type"`
	JSONData  string    `json:"json_data"`
}

type Gallery struct {
	GraphList []*Graph `json:"graph_list"`
	EdgeList  []*Edge  `json:"edge_list"`
}

func ParseJSON(jsonData string) (*Gallery, error) {
	galleryData := &Gallery{}
	if err := json.Unmarshal([]byte(jsonData), galleryData); err != nil {
		return nil, err
	}

	return galleryData, nil
}

func (g *Gallery) ToAllASTNode() ([]lubanAST.INode, error) {
	// todo build main func
	var allASTNodeList []lubanAST.INode

	for _, graphData := range g.GraphList {
		if graphData.GraphType == GraphProgram {
			programGraph, err := program.ParseJSON(graphData.JSONData)
			if err != nil {
				return nil, err
			}

			astNodeList, err := programGraph.ToAllASTNode()
			if err != nil {
				return nil, err
			}
			allASTNodeList = append(allASTNodeList, astNodeList...)
		}
	}

	return allASTNodeList, nil
}
