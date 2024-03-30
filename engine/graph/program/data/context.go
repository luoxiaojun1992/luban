package data

import (
	commonElementsFunction "github.com/luoxiaojun1992/luban/engine/elements/function"
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type Context struct {
	Caller      *commonElementsFunction.Caller    `json:"caller"`
	IsAsync     bool                              `json:"is_async"`
	InputVars   []string                          `json:"input_vars"`
	OutputVars  []string                          `json:"output_vars"`
	Params      []*commonElementsFunction.Param   `json:"params"`
	OutputTypes []*commonElementsVariable.VarType `json:"output_types"`
}

func (c *Context) HasCaller() bool {
	return c.Caller != nil
}
