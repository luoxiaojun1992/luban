package stmt

import goAST "go/ast"

type TwoBranch struct {
	Cond      *Expr
	MatchBody *Block
	ElseBody  *Block
}

func (tb *TwoBranch) ToGoASTStmt() (goAST.Stmt, error) {
	goASTIfStmt := &goAST.IfStmt{}

	if tb.Cond != nil {
		goASTCondExpr, err := tb.Cond.ToGoASTExpr()
		if err != nil {
			return nil, err
		}
		goASTIfStmt.Cond = goASTCondExpr
	}

	if tb.MatchBody != nil {
		goASTBodyBlockStmt, err := tb.MatchBody.ToGoASTBlockStmt()
		if err != nil {
			return nil, err
		}
		goASTIfStmt.Body = goASTBodyBlockStmt
	}

	if tb.ElseBody != nil {
		goASTElseBodyStmt, err := tb.ElseBody.ToGoASTStmt()
		if err != nil {
			return nil, err
		}
		goASTIfStmt.Else = goASTElseBodyStmt
	}

	return goASTIfStmt, nil
}

func (tb *TwoBranch) ToGoASTNode() (goAST.Node, error) {
	return tb.ToGoASTStmt()
}
