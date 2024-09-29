package types

type DriversData struct {
	Limit       string `json:"limit,omitempty"`
	Offset      string `json:"offset,omitempty"`
	Total       string `json:"total,omitempty"`
	Series      string `json:"series,omitempty"`
	Url         string `json:"url,omitempty"`
	Xmlns       string `json:"xmlns,omitempty"`
	DriverTable struct {
		Season   string `json:"season,omitempty"`
		Round    string `json:"round,omitempty"`
		DriverId string `json:"driverId,omitempty"`
		Drivers  []struct {
			DriverId        string `json:"driverId,omitempty"`
			PermanentNumber string `json:"permanentNumber,omitempty"`
			Code            string `json:"code,omitempty"`
			URL             string `json:"url,omitempty"`
			GivenName       string `json:"givenName,omitempty"`
			FamilyName      string `json:"FamilyName,omitempty"`
			DateOfBirth     string `json:"dateOfBirth,omitempty"`
			Nationality     string `json:"nationality,omitempty"`
		} `json:"Drivers,omitempty"`
	} `json:"DriverTable,omitempty"`
}

type Driver struct {
	DriverId        string `json:"driverId,omitempty"`
	PermanentNumber string `json:"permanentNumber,omitempty"`
	Code            string `json:"code,omitempty"`
	URL             string `json:"url,omitempty"`
	GivenName       string `json:"givenName,omitempty"`
	FamilyName      string `json:"FamilyName,omitempty"`
	DateOfBirth     string `json:"dateOfBirth,omitempty"`
	Nationality     string `json:"nationality,omitempty"`
}
