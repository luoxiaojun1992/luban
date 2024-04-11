package function

import (
	"go/ast"

	"github.com/luoxiaojun1992/luban/engine/ast/luban/stmt"
	commonElementsFunction "github.com/luoxiaojun1992/luban/engine/elements/function"
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type Function struct {
	Name        string
	Comments    []string
	Caller      *commonElementsFunction.Caller
	Params      []*commonElementsFunction.Param
	OutputTypes []*commonElementsVariable.VarType
	Body        *stmt.Block
}

func (f *Function) ToGoASTFuncDecl() (*ast.FuncDecl, error) {
	goASTFuncDecl := &ast.FuncDecl{}
	goASTFuncDecl.Name = ast.NewIdent(f.Name)

	if len(f.Comments) > 0 {
		goASTCommentGroup := &ast.CommentGroup{
			List: make([]*ast.Comment, 0, len(f.Comments)),
		}
		for _, comment := range f.Comments {
			goASTCommentGroup.List = append(goASTCommentGroup.List, &ast.Comment{
				Text: comment,
			})
		}

		goASTFuncDecl.Doc = goASTCommentGroup
	}

	if f.Caller != nil {
		goASTRecvFieldList := &ast.FieldList{
			List: []*ast.Field{{
				Names: []*ast.Ident{ast.NewIdent(f.Caller.Name)},
				Type:  ast.NewIdent(f.Caller.Type.String()),
			}},
		}
		goASTFuncDecl.Recv = goASTRecvFieldList
	}

	funcType := &ast.FuncType{}

	if len(f.Params) > 0 {
		goASTParamsFieldList := &ast.FieldList{
			List: make([]*ast.Field, 0, len(f.Params)),
		}
		for _, param := range f.Params {
			// todo support multi param names with same type
			goASTParamsFieldList.List = append(goASTParamsFieldList.List, &ast.Field{
				Names: []*ast.Ident{ast.NewIdent(param.Name)},
				Type:  ast.NewIdent(param.Type.String()),
			})
		}

		funcType.Params = goASTParamsFieldList
	}

	if len(f.OutputTypes) > 0 {
		goASTResultsFieldList := &ast.FieldList{
			List: make([]*ast.Field, 0, len(f.OutputTypes)),
		}
		for _, outputType := range f.OutputTypes {
			// todo support output var name
			goASTResultsFieldList.List = append(goASTResultsFieldList.List, &ast.Field{
				Type: ast.NewIdent(outputType.String()),
			})
		}

		funcType.Results = goASTResultsFieldList
	}

	goASTFuncDecl.Type = funcType

	goASTBlockStmt, err := f.Body.ToGoASTBlockStmt()
	if err != nil {
		return nil, err
	}
	goASTFuncDecl.Body = goASTBlockStmt

	return goASTFuncDecl, nil
}

func (f *Function) ToGoASTNode() (ast.Node, error) {
	return f.ToGoASTFuncDecl()
}
