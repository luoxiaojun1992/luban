package operand

type Raw struct {
	Raw string `json:"value"`
}

func (r *Raw) ToRaw() string {
	return r.Raw
}
