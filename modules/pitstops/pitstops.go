package pitstops

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

func PitstopsRequest(path string, query string) (types.PitstopsData, error) {
	resp, err := connection.Get(path, query)
	if err != nil {
		return types.PitstopsData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.PitstopsData{}, err
	}
	defer resp.Body.Close()

	var response types.PitstopsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return types.PitstopsData{}, err
	}
	return response.Data, nil
}

// http://ergast.com/api/f1/{year}/{round}/pitstops.json
func ByRace(year int, round int, offset int, limit int) (types.PitstopsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.PITSTOPS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return PitstopsRequest(path, query)
}

// http://ergast.com/api/f1/{year}/{round}/pitstops/{number}.json
func GetByNumber(year int, round int, number int, offset int, limit int) (types.PitstopsData, error) {
	path := fmt.Sprintf("%d/%d/%s/%d", year, round, urls.PITSTOPS, number)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return PitstopsRequest(path, query)
}
