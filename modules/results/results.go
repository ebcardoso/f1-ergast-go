package results

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

func ResultRequest(path string, query string) (types.ResultsData, error) {
	resp, err := connection.Get(path, query)
	if err != nil {
		return types.ResultsData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.ResultsData{}, err
	}
	defer resp.Body.Close()

	var response types.ResultsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return types.ResultsData{}, err
	}
	return response.Data, nil
}

// ergast.com/api/f1/{year}/{round}/results.json
func ByRace(year int, round int) (types.ResultsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.RESULT_RACE)
	return ResultRequest(path, "")
}

// ergast.com/api/f1/current/last/results.json
func MostRecent() (types.ResultsData, error) {
	return ResultRequest(urls.RESULT_MOST_RECENT, "")
}
