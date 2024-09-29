package types

type FinishingStatusData struct {
	Limit       string `json:"limit,omitempty"`
	Offset      string `json:"offset,omitempty"`
	Total       string `json:"total,omitempty"`
	Series      string `json:"series,omitempty"`
	Url         string `json:"url,omitempty"`
	Xmlns       string `json:"xmlns,omitempty"`
	StatusTable struct {
		Season string         `json:"season,omitempty"`
		Round  string         `json:"round,omitempty"`
		Status []FinishStatus `json:"Status,omitempty"`
	} `json:"StatusTable,omitempty"`
}

type FinishStatus struct {
	StatusId string `json:"statusId,omitempty"`
	Count    string `json:"count,omitempty"`
	Status   string `json:"status,omitempty"`
}
