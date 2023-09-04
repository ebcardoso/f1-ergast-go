package laps

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

func LapsRequest(path string, query string) (types.LapsData, error) {
	resp, err := connection.Get(path, query)
	if err != nil {
		return types.LapsData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.LapsData{}, err
	}
	defer resp.Body.Close()

	var response types.LapsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return types.LapsData{}, err
	}
	return response.Data, nil
}

// http://ergast.com/api/f1/{year}/{round}/laps/{lap}.json
func GetLapTimes(year int, round int, lap int, offset int, limit int) (types.LapsData, error) {
	path := fmt.Sprintf("%d/%d/%s/%d", year, round, urls.LAPS, lap)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return LapsRequest(path, query)
}
