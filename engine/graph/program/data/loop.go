package data

type Loop struct {
	Cond     string  `json:"cond"`
	CodeList []*Code `json:"code_list"`
}
