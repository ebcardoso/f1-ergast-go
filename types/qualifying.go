package types

type QualifyingData struct {
	paginationFields
	RaceTable QualifyingRaceTable `json:"RaceTable,omitempty"`
}

type QualifyingRaceTable struct {
	Season string            `json:"season,omitempty"`
	Round  string            `json:"round,omitempty"`
	Races  []QualifyingRaces `json:"Races,omitempty"`
}

type QualifyingRaces struct {
	Season            string              `json:"season,omitempty"`
	Round             string              `json:"round,omitempty"`
	Url               string              `json:"url,omitempty"`
	RaceName          string              `json:"raceName,omitempty"`
	Circuit           Circuit             `json:"Circuit,omitempty"`
	Date              string              `json:"date,omitempty"`
	Time              string              `json:"time,omitempty"`
	QualifyingResults []QualifyingResults `json:"QualifyingResults,omitempty"`
}

type QualifyingResults struct {
	Number      string      `json:"number,omitempty"`
	Position    string      `json:"position,omitempty"`
	Driver      Driver      `json:"Driver,omitempty"`
	Constructor Constructor `json:"Constructor,omitempty"`
	Q1          string      `json:"Q1,omitempty"`
	Q2          string      `json:"Q2,omitempty"`
	Q3          string      `json:"Q3,omitempty"`
}
