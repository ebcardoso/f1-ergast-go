package types

type DriverStandings struct {
	Position     string        `json:"position,omitempty"`
	PositionText string        `json:"positionText,omitempty"`
	Points       string        `json:"points,omitempty"`
	Wins         string        `json:"wins,omitempty"`
	Driver       Driver        `json:"Driver,omitempty"`
	Constructors []Constructor `json:"Constructors,omitempty"`
}

type ConstructorStandings struct {
	Position     string      `json:"position,omitempty"`
	PositionText string      `json:"positionText,omitempty"`
	Points       string      `json:"points,omitempty"`
	Wins         string      `json:"wins,omitempty"`
	Constructor  Constructor `json:"Constructor,omitempty"`
}

type StandingsData struct {
	Limit          string `json:"limit,omitempty"`
	Offset         string `json:"offset,omitempty"`
	Total          string `json:"total,omitempty"`
	Series         string `json:"series,omitempty"`
	Url            string `json:"url,omitempty"`
	Xmlns          string `json:"xmlns,omitempty"`
	StandingsTable struct {
		Season               string `json:"season,omitempty"`
		Round                string `json:"round,omitempty"`
		DriverId             string `json:"driverId,omitempty"`
		ConstructorId        string `json:"constructorId,omitempty"`
		DriverStandings      string `json:"driverStandings,omitempty"`
		ConstructorStandings string `json:"constructorStandings,omitempty"`
		StandingsLists       []struct {
			Season               string                 `json:"season,omitempty"`
			Round                string                 `json:"round,omitempty"`
			DriverStandings      []DriverStandings      `json:"DriverStandings,omitempty"`
			ConstructorStandings []ConstructorStandings `json:"ConstructorStandings,omitempty"`
		} `json:"StandingsLists,omitempty"`
	} `json:"StandingsTable,omitempty"`
}
