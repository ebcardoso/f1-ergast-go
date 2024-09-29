package types

type RacesData struct {
	Limit     string `json:"limit,omitempty"`
	Offset    string `json:"offset,omitempty"`
	Total     string `json:"total,omitempty"`
	Series    string `json:"series,omitempty"`
	Url       string `json:"url,omitempty"`
	Xmlns     string `json:"xmlns,omitempty"`
	RaceTable struct {
		Season string `json:"season,omitempty"`
		Round  string `json:"round,omitempty"`
		Races  []Race `json:"Races,omitempty"`
	} `json:"RaceTable"`
}

type Race struct {
	Season        string  `json:"season,omitempty"`
	Round         string  `json:"round,omitempty"`
	Url           string  `json:"url,omitempty"`
	RaceName      string  `json:"raceName,omitempty"`
	Circuit       Circuit `json:"Circuit,omitempty"`
	Date          string  `json:"date,omitempty"`
	Time          string  `json:"time,omitempty"`
	FirstPractice struct {
		Date string `json:"date,omitempty"`
		Time string `json:"time,omitempty"`
	} `json:"FirstPractice,omitempty"`
	SecondPractice struct {
		Date string `json:"date,omitempty"`
		Time string `json:"time,omitempty"`
	} `json:"SecondPractice,omitempty"`
	ThirdPractice struct {
		Date string `json:"date,omitempty"`
		Time string `json:"time,omitempty"`
	} `json:"ThirdPractice,omitempty"`
	Qualifying struct {
		Date string `json:"date,omitempty"`
		Time string `json:"time,omitempty"`
	} `json:"Qualifying,omitempty"`
}
