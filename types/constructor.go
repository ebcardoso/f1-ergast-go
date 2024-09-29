package types

type ConstructorsData struct {
	paginationFields
	ConstructorTable struct {
		Season        string        `json:"season,omitempty"`
		Round         string        `json:"round,omitempty"`
		ConstructorId string        `json:"constructorId,omitempty"`
		Constructors  []Constructor `json:"Constructors,omitempty"`
	} `json:"ConstructorTable,omitempty"`
}

type Constructor struct {
	ConstructorId string `json:"constructorId,omitempty"`
	URL           string `json:"url,omitempty"`
	Name          string `json:"name,omitempty"`
	Nationality   string `json:"nationality,omitempty"`
}
