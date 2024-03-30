package function

import "github.com/luoxiaojun1992/luban/engine/elements/variable"

type Param struct {
	Name string            `json:"name"`
	Type *variable.VarType `json:"type"`
}

func (p *Param) Copy() *Param {
	return &Param{
		Name: p.Name,
		Type: p.Type.Copy(),
	}
}
