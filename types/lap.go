package types

type LapsData struct {
	Limit     string `json:"limit,omitempty"`
	Offset    string `json:"offset,omitempty"`
	Total     string `json:"total,omitempty"`
	Series    string `json:"series,omitempty"`
	Url       string `json:"url,omitempty"`
	Xmlns     string `json:"xmlns,omitempty"`
	RaceTable struct {
		Season string `json:"season,omitempty"`
		Round  string `json:"round,omitempty"`
		Races  []struct {
			Season   string  `json:"season,omitempty"`
			Round    string  `json:"round,omitempty"`
			Url      string  `json:"url,omitempty"`
			RaceName string  `json:"raceName,omitempty"`
			Circuit  Circuit `json:"Circuit,omitempty"`
			Date     string  `json:"date,omitempty"`
			Time     string  `json:"time,omitempty"`
			Laps     []Lap   `json:"Laps,omitempty"`
		} `json:"Races,omitempty"`
	} `json:"RaceTable,omitempty"`
}

type Lap struct {
	Number  string   `json:"number,omitempty"`
	Timings []Timing `json:"Timings,omitempty"`
}

type Timing struct {
	DriverId string `json:"driverId,omitempty"`
	Position string `json:"position,omitempty"`
	Time     string `json:"time,omitempty"`
}
