//go:build js && wasm
// +build js,wasm

package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"go/token"
	"syscall/js"

	"github.com/luoxiaojun1992/luban/engine/graph/program"
)

func LubanProgramParse(_ js.Value, inputs []js.Value) interface{} {
	if len(inputs) != 1 {
		panic("invalid input")
	}

	jsonData := inputs[0].String()

	graphModel, err := program.ParseJSON(jsonData)
	if err != nil {
		panic(err)
	}

	lubanASTNodeList, err := graphModel.ToAllASTNode()
	if err != nil {
		panic(err)
	}

	result := make([]string, 0, len(lubanASTNodeList))

	for _, lubanASTNode := range lubanASTNodeList {
		goASTNode, err := lubanASTNode.ToGoASTNode()
		if err != nil {
			panic(err)
		}

		buf := bytes.NewBuffer(nil)
		if err := format.Node(buf, token.NewFileSet(), goASTNode); err != nil {
			panic(err)
		}

		result = append(result, buf.String())
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return js.ValueOf(string(resultJson))
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("LubanProgramParse", js.FuncOf(LubanProgramParse))
	<-done
}
