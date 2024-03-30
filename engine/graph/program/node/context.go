package node

import (
	commonElementsFunction "github.com/luoxiaojun1992/luban/engine/elements/function"
	commonElementsVariable "github.com/luoxiaojun1992/luban/engine/elements/variable"
)

type Context struct {
	Caller      *commonElementsFunction.Caller
	IsAsync     bool
	InputVars   []string
	OutputVars  []string
	Params      []*commonElementsFunction.Param
	OutputTypes []*commonElementsVariable.VarType

	Prev       INode
	Next       INode
	MatchChild INode
	ElseChild  INode
}

func (c *Context) HasCaller() bool {
	return c.Caller != nil
}

func (c *Context) GetCaller() *commonElementsFunction.Caller {
	return c.Caller
}

func (c *Context) CheckIsAsync() bool {
	return c.IsAsync
}

func (c *Context) HasInputVars() bool {
	return len(c.InputVars) > 0
}

func (c *Context) GetInputVars() []string {
	return c.InputVars
}

func (c *Context) HasOutputVars() bool {
	return len(c.OutputVars) > 0
}

func (c *Context) GetOutputVars() []string {
	return c.OutputVars
}

func (c *Context) GetParams() []*commonElementsFunction.Param {
	return c.Params
}

func (c *Context) GetOutputTypes() []*commonElementsVariable.VarType {
	return c.OutputTypes
}

func (c *Context) HasPrev() bool {
	return c.Prev != nil
}

func (c *Context) GetPrev() INode {
	return c.Prev
}

func (c *Context) SetPrev(prev INode) {
	c.Prev = prev
}

func (c *Context) HasNext() bool {
	return c.Next != nil
}

func (c *Context) GetNext() INode {
	return c.Next
}

func (c *Context) SetNext(next INode) {
	c.Next = next
}

func (c *Context) HasMatchChild() bool {
	return c.MatchChild != nil
}

func (c *Context) GetMatchChild() INode {
	return c.MatchChild
}

func (c *Context) SetMatchChild(matchChild INode) {
	c.MatchChild = matchChild
}

func (c *Context) HasElseChild() bool {
	return c.ElseChild != nil
}

func (c *Context) GetElseChild() INode {
	return c.ElseChild
}

func (c *Context) SetElseChild(elseChild INode) {
	c.ElseChild = elseChild
}
