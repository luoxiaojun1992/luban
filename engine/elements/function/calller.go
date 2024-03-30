package function

import "github.com/luoxiaojun1992/luban/engine/elements/variable"

type Caller struct {
	Name    string            `json:"name"`
	VarName string            `json:"var_name"`
	Type    *variable.VarType `json:"type"`
}

func (c *Caller) Copy() *Caller {
	return &Caller{
		Name:    c.Name,
		VarName: c.VarName,
		Type:    c.Type.Copy(),
	}
}
