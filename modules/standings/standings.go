package standings

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils"
)

func StandingsRequest(path string, query string) (types.StandingsData, error) {
	resp, err := connection.Get(path, query)
	if err != nil {
		return types.StandingsData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.StandingsData{}, err
	}
	defer resp.Body.Close()

	var response types.StandingsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return types.StandingsData{}, err
	}
	return response.Data, nil
}

// ergast.com/api/f1/{{year}}/{{round}}/driverStandings.json
func DriversAfterRace(year int, round int, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, utils.STANDINGS_DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return StandingsRequest(path, query)
}

// ergast.com/api/f1/{{year}}/{{round}}/constructorStandings.json
func ConstructorsAfterRace(year int, round int, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, utils.STANDINGS_CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return StandingsRequest(path, query)
}

// ergast.com/api/f1/{{year}}/driverStandings.json
func DriversBySeason(year int, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%d/%s", year, utils.STANDINGS_DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return StandingsRequest(path, query)
}

// ergast.com/api/f1/{{year}}/constructorStandings.json
func ConstructorsBySeason(year int, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%d/%s", year, utils.STANDINGS_CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return StandingsRequest(path, query)
}
