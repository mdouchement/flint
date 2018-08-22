package formatter

type summary struct {
	Errors   severitySummary `json:"errors"`
	Warnings severitySummary `json:"warnings"`
}

type severitySummary struct {
	Total uint64            `json:"total"`
	Rules map[string]uint64 `json:"rules"`
}
