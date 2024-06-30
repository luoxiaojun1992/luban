package node

type Table struct {
	Common
	Context
	Table string `json:"table"`
}

func (t *Table) GetTable() string {
	return t.Table
}

func (t *Table) GetField() *Field {
	return nil
}

func (t *Table) GetAssignment() *Assign {
	return nil
}

func (t *Table) GetCondition() *Condition {
	return nil
}
