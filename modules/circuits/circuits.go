package circuits

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

func CircuitsRequest(path string, query string) (types.CircuitsData, error) {
	resp, err := connection.Get(path, query)
	if err != nil {
		return types.CircuitsData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.CircuitsData{}, err
	}
	defer resp.Body.Close()

	var response types.CircuitsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return types.CircuitsData{}, err
	}
	return response.Data, nil
}

// ergast.com/api/f1/circuits.json
func List(offset int, limit int) (types.CircuitsData, error) {
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return CircuitsRequest(urls.CIRCUITS, query)
}

// ergast.com/api/f1/{year}/circuits.json
func BySeason(year int, offset int, limit int) (types.CircuitsData, error) {
	path := fmt.Sprintf("%d/%s", year, urls.CIRCUITS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return CircuitsRequest(path, query)
}

// ergast.com/api/f1/{year}/{round}/circuits.json
func ByRace(year int, round int, offset int, limit int) (types.CircuitsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.CIRCUITS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return CircuitsRequest(path, query)
}

// ergast.com/api/f1/{circuitId}/circuits.json
func GetByCircuitId(circuitId string) (types.CircuitsData, error) {
	path := fmt.Sprintf("%s/%s", urls.CIRCUITS, circuitId)

	return CircuitsRequest(path, "")
}
