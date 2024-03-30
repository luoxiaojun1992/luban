package stmt

import goAST "go/ast"

type MultiBranch struct {
	Needle  *Expr
	Matches []*Matcher
}

func (mb *MultiBranch) ToGoASTStmt() (goAST.Stmt, error) {
	goASTSwitchStmt := &goAST.SwitchStmt{}

	if mb.Needle != nil {
		goASTCondExpr, err := mb.Needle.ToGoASTExpr()
		if err != nil {
			return nil, err
		}
		goASTSwitchStmt.Tag = goASTCondExpr
	}

	if len(mb.Matches) > 0 {
		lubanASTBodyBlock := &Block{}
		for _, match := range mb.Matches {
			lubanASTBodyBlock.Stmts = append(lubanASTBodyBlock.Stmts, match)
		}

		goASTBodyBlockStmt, err := lubanASTBodyBlock.ToGoASTBlockStmt()
		if err != nil {
			return nil, err
		}

		goASTSwitchStmt.Body = goASTBodyBlockStmt
	}

	return goASTSwitchStmt, nil
}

func (mb *MultiBranch) ToGoASTNode() (goAST.Node, error) {
	return mb.ToGoASTStmt()
}
