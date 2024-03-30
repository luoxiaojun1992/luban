package stmt

import goAST "go/ast"

type Block struct {
	Stmts []IStmt
}

func (b *Block) ToGoASTBlockStmt() (*goAST.BlockStmt, error) {
	goAstBlockStmt := &goAST.BlockStmt{}

	for _, stmt := range b.Stmts {
		goASTStmt, err := stmt.ToGoASTStmt()
		if err != nil {
			return nil, err
		}
		goAstBlockStmt.List = append(goAstBlockStmt.List, goASTStmt)
	}

	return goAstBlockStmt, nil
}

func (b *Block) ToGoASTStmt() (goAST.Stmt, error) {
	return b.ToGoASTBlockStmt()
}

func (b *Block) ToGoASTNode() (goAST.Node, error) {
	return b.ToGoASTStmt()
}
