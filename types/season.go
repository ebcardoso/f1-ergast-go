package types

type SeasonsData struct {
	paginationFields
	SeasonTable struct {
		Seasons []Season `json:"Seasons,omitempty"`
	} `json:"SeasonTable,omitempty"`
}

type Season struct {
	Season string `json:"season,omitempty"`
	URL    string `json:"url,omitempty"`
}
