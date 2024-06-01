package stmt_test

import (
	goAST "go/ast"
	"testing"

	lubanASTStmt "github.com/luoxiaojun1992/luban/engine/ast/luban/stmt"
)

// TestBlockToGoASTBlockStmt tests the ToGoASTBlockStmt method of the Block struct
func TestBlockToGoASTBlockStmt(t *testing.T) {
	// Create a sample Block with some IStmt elements
	block := &lubanASTStmt.Block{
		Stmts: []lubanASTStmt.IStmt{
			&lubanASTStmt.Expr{Raw: "println(\"foo\")"},
			&lubanASTStmt.Expr{Raw: "println(\"bar\")"},
			&lubanASTStmt.Expr{Raw: "println(\"baz\")"},
		},
	}

	// Call the ToGoASTBlockStmt method
	goAstBlockStmt, err := block.ToGoASTBlockStmt()
	// Verify the returned values
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(goAstBlockStmt.List) != len(block.Stmts) {
		t.Errorf("Unexpected length of goAstBlockStmt.List. Expected: %d, Got: %d", len(block.Stmts), len(goAstBlockStmt.List))
	}
}

// TestBlockToGoASTStmt tests the ToGoASTStmt method of the Block struct
func TestBlockToGoASTStmt(t *testing.T) {
	// Create a sample Block
	block := &lubanASTStmt.Block{}

	// Call the ToGoASTStmt method
	goASTStmt, err := block.ToGoASTStmt()
	// Verify the returned values
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	goAstBlockStmt, ok := goASTStmt.(*goAST.BlockStmt)
	if !ok {
		t.Errorf("Unexpected type of goASTStmt. Expected: *goAST.BlockStmt, Got: %T", goASTStmt)
	}

	if goAstBlockStmt == nil {
		t.Errorf("Unexpected nil value for goAstBlockStmt")
	}
}

// TestBlockToGoASTNode tests the ToGoASTNode method of the Block struct
func TestBlockToGoASTNode(t *testing.T) {
	// Create a sample Block
	block := &lubanASTStmt.Block{}

	// Call the ToGoASTNode method
	goASTNode, err := block.ToGoASTNode()
	// Verify the returned values
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	goAstBlockStmt, ok := goASTNode.(*goAST.BlockStmt)
	if !ok {
		t.Errorf("Unexpected type of goASTNode. Expected: *goAST.BlockStmt, Got: %T", goASTNode)
	}

	if goAstBlockStmt == nil {
		t.Errorf("Unexpected nil value for goAstBlockStmt")
	}
}
