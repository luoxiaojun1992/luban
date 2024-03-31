package program_test

import (
	"bytes"
	"go/format"
	"go/token"
	"io"
	"os"
	"testing"

	"github.com/luoxiaojun1992/luban/engine/graph/program"
	"github.com/onsi/gomega"
)

func TestToAllASTNode(t *testing.T) {
	f, err := os.Open("../../../data/test_stubs/sample_graph.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	jsonBytes, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	graphModel, err := program.ParseJSON(string(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	astNodeList, err := graphModel.ToAllASTNode()
	if err != nil {
		t.Fatal(err)
	}

	assert := gomega.NewWithT(t)

	for _, astNode := range astNodeList {
		goASTNode, err := astNode.ToGoASTNode()
		if err != nil {
			t.Fatal(err)
		}

		buf := bytes.NewBuffer(nil)
		if err := format.Node(buf, token.NewFileSet(), goASTNode); err != nil {
			t.Fatal(err)
		}

		t.Log(buf.String())
		assert.Expect(buf.String()).NotTo(gomega.BeZero())
	}
}
