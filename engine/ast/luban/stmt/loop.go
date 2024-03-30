package stmt

import goAST "go/ast"

type Loop struct {
	Cond *Expr
	Body *Block
}

func (l *Loop) ToGoASTStmt() (goAST.Stmt, error) {
	goASTForStmt := &goAST.ForStmt{}

	if l.Cond != nil {
		goASTCondExpr, err := l.Cond.ToGoASTExpr()
		if err != nil {
			return nil, err
		}
		goASTForStmt.Cond = goASTCondExpr
	}

	if l.Body != nil {
		goASTBodyBlockStmt, err := l.Body.ToGoASTBlockStmt()
		if err != nil {
			return nil, err
		}
		goASTForStmt.Body = goASTBodyBlockStmt
	}

	return goASTForStmt, nil
}

func (l *Loop) ToGoASTNode() (goAST.Node, error) {
	return l.ToGoASTStmt()
}
