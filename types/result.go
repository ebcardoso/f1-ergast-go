package types

type ResultsData struct {
	paginationFields
	RaceTable RaceTable `json:"RaceTable,omitempty"`
}

type RaceTable struct {
	Season string       `json:"season,omitempty"`
	Round  string       `json:"round,omitempty"`
	Races  []RaceResult `json:"Races,omitempty"`
}

type RaceResult struct {
	Season   string   `json:"season,omitempty"`
	Round    string   `json:"round,omitempty"`
	Url      string   `json:"url,omitempty"`
	RaceName string   `json:"raceName,omitempty"`
	Circuit  Circuit  `json:"Circuit,omitempty"`
	Date     string   `json:"date,omitempty"`
	Time     string   `json:"time,omitempty"`
	Results  []Result `json:"Results,omitempty"`
}

type Result struct {
	Number       string      `json:"number,omitempty"`
	Position     string      `json:"position,omitempty"`
	PositionText string      `json:"positionText,omitempty"`
	Points       string      `json:"points,omitempty"`
	Driver       Driver      `json:"Driver,omitempty"`
	Constructor  Constructor `json:"Constructor,omitempty"`
	Grid         string      `json:"grid,omitempty"`
	Laps         string      `json:"laps,omitempty"`
	Status       string      `json:"status,omitempty"`
	Time         struct {
		Millis string `json:"millis,omitempty"`
		Time   string `json:"time,omitempty"`
	} `json:"Time,omitempty"`
	FastestLap FastestLap `json:"FastestLap,omitempty"`
}
