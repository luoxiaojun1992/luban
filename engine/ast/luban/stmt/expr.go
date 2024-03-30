package stmt

import (
	"errors"
	goAST "go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type Expr struct {
	Raw          string
	IsBreakStmt  bool
	IsRetrunStmt bool
	IsAsyncStmt  bool

	BreakStmt  *BreakStmt
	ReturnStmt *ReturnStmt
	AsyncStmt  *AsyncStmt
}

func (e *Expr) AssertStmtType() *Expr {
	e.AssertBreakStmt()
	e.AssertReturnStmt()
	e.AssertAsyncStmt()
	return e
}

func (e *Expr) AssertBreakStmt() *Expr {
	breakLabel := ""

	raw := strings.TrimLeft(e.Raw, " ")

	breakStmtPrefix := "break "
	if strings.HasPrefix(raw, breakStmtPrefix) {
		e.IsBreakStmt = true
		breakLabel = strings.TrimSpace(strings.TrimPrefix(raw, breakStmtPrefix))
	} else if raw == "break" {
		e.IsBreakStmt = true
	}

	if e.IsBreakStmt {
		e.BreakStmt = &BreakStmt{
			Label: breakLabel,
		}
	}

	return e
}

func (e *Expr) AssertReturnStmt() *Expr {
	returnExpr := ""

	raw := strings.TrimLeft(e.Raw, " ")

	returnStmtPrefix := "return "
	if strings.HasPrefix(raw, returnStmtPrefix) {
		e.IsRetrunStmt = true
		returnExpr = strings.TrimSpace(strings.TrimPrefix(raw, returnStmtPrefix))
	} else if raw == "return" {
		e.IsRetrunStmt = true
	}

	if e.IsRetrunStmt {
		e.ReturnStmt = &ReturnStmt{
			Expr: returnExpr,
		}
	}

	return e
}

func (e *Expr) AssertAsyncStmt() *Expr {
	callExpr := ""

	raw := strings.TrimLeft(e.Raw, " ")

	goStmtPrefix := "go "
	if strings.HasPrefix(raw, goStmtPrefix) {
		e.IsAsyncStmt = true
		callExpr = strings.TrimSpace(strings.TrimPrefix(raw, goStmtPrefix))
	}

	if e.IsAsyncStmt {
		e.AsyncStmt = &AsyncStmt{
			CallExpr: callExpr,
		}
	}

	return e
}

func (e *Expr) ToGoASTExpr() (goAST.Expr, error) {
	if e.IsBreakStmt {
		return nil, errors.New("break stmt is not supported")
	}

	raw := e.Raw

	if e.IsRetrunStmt {
		if len(e.ReturnStmt.Expr) <= 0 {
			return nil, errors.New("return stmt expr not found")
		}

		raw = e.ReturnStmt.Expr
	} else if e.IsAsyncStmt {
		raw = e.AsyncStmt.CallExpr
	}

	return parser.ParseExpr(raw)
}

func (e *Expr) ToGoASTStmt() (goAST.Stmt, error) {
	if e.IsBreakStmt {
		goASTBranchStmt := &goAST.BranchStmt{
			Tok: token.BREAK,
		}
		if len(e.BreakStmt.Label) > 0 {
			goASTBranchStmt.Label = goAST.NewIdent(e.BreakStmt.Label)
		}
		return goASTBranchStmt, nil
	}

	if e.IsRetrunStmt {
		goASTReturnStmt := &goAST.ReturnStmt{}
		if len(e.ReturnStmt.Expr) > 0 {
			goASTExpr, err := e.ToGoASTExpr()
			if err != nil {
				return nil, err
			}
			goASTReturnStmt.Results = append(goASTReturnStmt.Results, goASTExpr)
		}
		return goASTReturnStmt, nil
	}

	if e.IsAsyncStmt {
		goASTExpr, err := e.ToGoASTExpr()
		if err != nil {
			return nil, err
		}

		if goASTCallExpr, ok := goASTExpr.(*goAST.CallExpr); ok {
			return &goAST.GoStmt{
				Call: goASTCallExpr,
			}, nil
		}
		return nil, errors.New("invalid async stmt call expr")
	}

	goASTExpr, err := e.ToGoASTExpr()
	if err != nil {
		return nil, err
	}

	return &goAST.ExprStmt{
		X: goASTExpr,
	}, nil
}

func (e *Expr) ToGoASTNode() (goAST.Node, error) {
	return e.ToGoASTStmt()
}
