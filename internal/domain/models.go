package domain

type Metadata struct {
	TotalCount int64 `json:"total_count"`
	PageCount  int64 `json:"page_count"`
	PageLimit  int64 `json:"page_limit"`
	Skipped    int64 `json:"skipped"`
	Sorted     int64 `json:"sorted"`
}

type Sort struct {
	Field string `json:"field"`
	Order int64  `json:"order"`
}
