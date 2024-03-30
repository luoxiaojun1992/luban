package sql

type IOperand interface {
	ToRaw() string
}
