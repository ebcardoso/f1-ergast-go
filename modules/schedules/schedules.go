package schedules

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils"
)

func SchedulesRequest(path string, query string) (types.RacesData, error) {
	resp, err := connection.Get(path, query)
	if err != nil {
		return types.RacesData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.RacesData{}, err
	}
	defer resp.Body.Close()

	var response types.RacesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return types.RacesData{}, err
	}
	return response.Data, nil
}

// ergast.com/api/f1/{year}.json
func BySeason(year int, offset int, limit int) (types.RacesData, error) {
	path := fmt.Sprintf("%d", year)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return SchedulesRequest(path, query)
}

// ergast.com/api/f1/current.json
func ByCurrentSeason(offset int, limit int) (types.RacesData, error) {
	path := utils.SCHEDULE_CURRENT
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return SchedulesRequest(path, query)
}

// ergast.com/api/f1/{year}/{round}.json
func ByRace(year int, round int) (types.RacesData, error) {
	path := fmt.Sprintf("%d/%d", year, round)
	query := ""

	return SchedulesRequest(path, query)
}
