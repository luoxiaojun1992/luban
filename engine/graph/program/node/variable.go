package node

type Variable struct {
	Name string
	Lock *int
}

func NewVariable(name string) *Variable {
	return &Variable{
		Name: name,
	}
}

func (nv *Variable) TryRLock() bool {
	if nv.Lock != nil && *nv.Lock == 1 {
		return false
	}

	return true
}

func (nv *Variable) RLock() bool {
	if !nv.TryRLock() {
		return false
	}

	rLock := 0
	nv.Lock = &rLock
	return true
}

func (nv *Variable) TryExLock() bool {
	if nv.Lock != nil && *nv.Lock|1 == 1 {
		return false
	}

	return true
}

func (nv *Variable) ExLock() bool {
	if !nv.TryExLock() {
		return false
	}

	exLock := 1
	nv.Lock = &exLock
	return true
}
