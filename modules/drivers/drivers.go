package drivers

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

func DriversRequest(path string, query string) (types.DriversData, error) {
	resp, err := connection.Get(path, query)
	if err != nil {
		return types.DriversData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.DriversData{}, err
	}
	defer resp.Body.Close()

	var response types.DriversResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return types.DriversData{}, err
	}
	return response.Data, nil
}

// ergast.com/api/f1/drivers.json
func List(offset int, limit int) (types.DriversData, error) {
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return DriversRequest(urls.DRIVERS, query)
}

// ergast.com/api/f1/{year}/drivers.json
func BySeason(year int, offset int, limit int) (types.DriversData, error) {
	path := fmt.Sprintf("%d/%s", year, urls.DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return DriversRequest(path, query)
}

// ergast.com/api/f1/{year}/{round}/drivers.json
func ByRace(year int, round int, offset int, limit int) (types.DriversData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return DriversRequest(path, query)
}

// ergast.com/api/f1/{driverId}/drivers.json
func GetByDriverId(driverId string) (types.DriversData, error) {
	path := fmt.Sprintf("%s/%s", urls.DRIVERS, driverId)

	return DriversRequest(path, "")
}
