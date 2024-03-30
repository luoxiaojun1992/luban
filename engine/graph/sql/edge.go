package sql

type Edge struct {
	Src   int                    `json:"src"`
	Dst   int                    `json:"dst"`
	Attrs map[string]interface{} `json:"attrs"`
}
