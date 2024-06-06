package node

type Table struct {
	Common
	Context
	Table string `json:"table"`
}

func (t *Table) GetTable() string {
	return t.Table
}
