package types

type FinishingStatusData struct {
	paginationFields
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
