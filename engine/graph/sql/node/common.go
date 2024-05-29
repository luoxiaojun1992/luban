package node

type Common struct {
	ID       int
	Name     string
	NodeType NodeType
}

func (c *Common) GetID() int {
	return c.ID
}

func (c *Common) GetName() string {
	return c.Name
}

func (c *Common) GetType() NodeType {
	return c.NodeType
}
