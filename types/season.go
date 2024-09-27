package types

type Season struct {
	Season string `json:"season,omitempty"`
	URL    string `json:"url,omitempty"`
}

type SeasonsData struct {
	Limit       string `json:"limit,omitempty"`
	Offset      string `json:"offset,omitempty"`
	Total       string `json:"total,omitempty"`
	Series      string `json:"series,omitempty"`
	Url         string `json:"url,omitempty"`
	Xmlns       string `json:"xmlns,omitempty"`
	SeasonTable struct {
		Seasons []Season `json:"Seasons,omitempty"`
	} `json:"SeasonTable,omitempty"`
}
