package graph

import (
	"encoding/json"

	lubanAST "github.com/luoxiaojun1992/luban/engine/ast/luban"
)

type GraphType string

const (
	GraphProgram GraphType = "program"
	GraphSQL     GraphType = "sql"
)

type IGraph interface {
	ToAllASTNode() ([]lubanAST.INode, error)
}

type Gallery struct {
	GraphList []string `json:"graph_list"`
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
	//todo
	return nil, nil
}
