package node

type Table struct {
	Common
	Context
	Name string `json:"name"`
}

func (t *Table) GetTable() string {
	return t.Name
}
