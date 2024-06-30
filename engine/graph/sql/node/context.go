package node

type Context struct {
	Prev INode
	Next INode
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
