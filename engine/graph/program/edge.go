package program

type Edge struct {
	Src   int                    `json:"src"`
	Dst   int                    `json:"dst"`
	Attrs map[string]interface{} `json:"attrs"`
}

func (e *Edge) GetBranchType() string {
	if branchType, ok := e.Attrs["branch_type"]; ok {
		if branchTypeStr, isStr := branchType.(string); isStr {
			return branchTypeStr
		}
	}

	return ""
}
