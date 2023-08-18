package types

type FastestLap struct {
	Rank string `json:"rank,omitempty"`
	Lap  string `json:"lap,omitempty"`
	Time struct {
		Time string `json:"time,omitempty"`
	} `json:"Time,omitempty"`
	AverageSpeed struct {
		Units string `json:"units,omitempty"`
		Speed string `json:"speed,omitempty"`
	} `json:"AverageSpeed,omitempty"`
}
