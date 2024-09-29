package types

type CircuitsData struct {
	paginationFields
	CircuitTable struct {
		Season    string    `json:"season,omitempty"`
		Round     string    `json:"round,omitempty"`
		CircuitId string    `json:"circuitId,omitempty"`
		Circuits  []Circuit `json:"Circuits,omitempty"`
	} `json:"CircuitTable,omitempty"`
}

type Circuit struct {
	CircuitId   string `json:"circuitId,omitempty"`
	Url         string `json:"url,omitempty"`
	CircuitName string `json:"circuitName,omitempty"`
	Location    struct {
		Lat      string `json:"lat,omitempty"`
		Long     string `json:"long,omitempty"`
		Locality string `json:"locality,omitempty"`
		Country  string `json:"country,omitempty"`
	} `json:"location,omitempty"`
}
