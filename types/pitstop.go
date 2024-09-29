package types

type PitstopsData struct {
	paginationFields
	RaceTable struct {
		Season string `json:"season,omitempty"`
		Round  string `json:"round,omitempty"`
		Stop   string `json:"stop,omitempty"`
		Races  []struct {
			Season   string    `json:"season,omitempty"`
			Round    string    `json:"round,omitempty"`
			Url      string    `json:"url,omitempty"`
			RaceName string    `json:"racename,omitempty"`
			Circuit  Circuit   `json:"Circuit,omitempty"`
			Date     string    `json:"date,omitempty"`
			Time     string    `json:"time,omitempty"`
			Pitstops []Pitstop `json:"PitStops,omitempty"`
		} `json:"Races,omitempty"`
	} `json:"RaceTable,omitempty"`
}

type Pitstop struct {
	DriverId string `json:"driverId,omitempty"`
	Lap      string `json:"lap,omitempty"`
	Stop     string `json:"stop,omitempty"`
	Time     string `json:"time,omitempty"`
	Duration string `json:"duration,omitempty"`
}
