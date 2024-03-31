package variable

import "fmt"

type Value struct {
	Value   string   `json:"value"`
	ValType *VarType `json:"val_type"`
}

func (v *Value) ToString() string {
	value := v.Value
	if v.ValType.IsString() {
		value = fmt.Sprintf("\"%s\"", v.Value)
	}
	return value
}
