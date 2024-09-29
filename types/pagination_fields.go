package types

type paginationFields struct {
	Limit  string `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
	Total  string `json:"total,omitempty"`
	Series string `json:"series,omitempty"`
	Url    string `json:"url,omitempty"`
	Xmlns  string `json:"xmlns,omitempty"`
}
