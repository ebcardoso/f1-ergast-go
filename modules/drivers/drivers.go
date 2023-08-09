package drivers

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/utils"
)

type DriversResponse struct {
	MRData struct {
		Limit       string `json:"limit,omitempty"`
		Offset      string `json:"offset,omitempty"`
		Total       string `json:"total,omitempty"`
		Series      string `json:"series,omitempty"`
		Url         string `json:"url,omitempty"`
		Xmlns       string `json:"xmlns,omitempty"`
		DriverTable struct {
			Season   string `json:"season,omitempty"`
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
			}
		} `json:"DriverTable,omitempty"`
	} `json:"MRData,omitempty"`
}

func DriversRequest(path string, query string) (DriversResponse, error) {
	resp, err := connection.Get(path, query)
	if err != nil {
		return DriversResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return DriversResponse{}, err
	}
	defer resp.Body.Close()

	var response DriversResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return DriversResponse{}, err
	}
	return response, nil
}

// ergast.com/api/f1/drivers.json
func List(offset int, limit int) (DriversResponse, error) {
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return DriversRequest(utils.DRIVERS, query)
}

// ergast.com/api/f1/{year}/drivers.json
func BySeason(year int, offset int, limit int) (DriversResponse, error) {
	path := fmt.Sprintf("%d/%s", year, utils.DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return DriversRequest(path, query)
}

// ergast.com/api/f1/{year}/{round}/drivers.json
func ByRace(year int, round int, offset int, limit int) (DriversResponse, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, utils.DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return DriversRequest(path, query)
}

// ergast.com/api/f1/{driverId}/drivers.json
func GetByDriverId(driverId string) (DriversResponse, error) {
	path := fmt.Sprintf("%s/%s", utils.DRIVERS, driverId)

	return DriversRequest(path, "")
}
