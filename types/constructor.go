package types

type Constructor struct {
	ConstructorId string `json:"constructorId,omitempty"`
	URL           string `json:"url,omitempty"`
	Name          string `json:"name,omitempty"`
	Nationality   string `json:"nationality,omitempty"`
}

type ConstructorsData struct {
	Limit            string `json:"limit,omitempty"`
	Offset           string `json:"offset,omitempty"`
	Total            string `json:"total,omitempty"`
	Series           string `json:"series,omitempty"`
	Url              string `json:"url,omitempty"`
	Xmlns            string `json:"xmlns,omitempty"`
	ConstructorTable struct {
		Season        string        `json:"season,omitempty"`
		Round         string        `json:"round,omitempty"`
		ConstructorId string        `json:"constructorId,omitempty"`
		Constructors  []Constructor `json:"Constructors,omitempty"`
	} `json:"ConstructorTable,omitempty"`
}
