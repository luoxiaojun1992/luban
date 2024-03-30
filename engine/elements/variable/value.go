package variable

type Value struct {
	Value   string   `json:"value"`
	ValType *VarType `json:"val_type"`
}
