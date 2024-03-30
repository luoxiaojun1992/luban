package stmt

import goAST "go/ast"

type Matcher struct {
	Targets []*Expr
	Body    *Block
}

func (m *Matcher) ToGoASTStmt() (goAST.Stmt, error) {
	goASTCaseClause := &goAST.CaseClause{}

	for _, target := range m.Targets {
		goASTTargetExpr, err := target.ToGoASTExpr()
		if err != nil {
			return nil, err
		}
		goASTCaseClause.List = append(goASTCaseClause.List, goASTTargetExpr)
	}

	if m.Body != nil {
		for _, bodyStmt := range m.Body.Stmts {
			goASTBodyStmt, err := bodyStmt.ToGoASTStmt()
			if err != nil {
				return nil, err
			}
			goASTCaseClause.Body = append(goASTCaseClause.Body, goASTBodyStmt)
		}
	}

	return goASTCaseClause, nil
}

func (m *Matcher) ToGoASTNode() (goAST.Node, error) {
	return m.ToGoASTStmt()
}
