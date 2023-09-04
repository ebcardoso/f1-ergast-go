package qualifying

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

func QualifyingRequest(path string, query string) (types.QualifyingData, error) {
	resp, err := connection.Get(path, query)
	if err != nil {
		return types.QualifyingData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.QualifyingData{}, err
	}
	defer resp.Body.Close()

	var response types.QualifyingResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return types.QualifyingData{}, err
	}
	return response.Data, nil
}

// ergast.com/api/f1/{year}/{round}/qualifying.json
func ByRace(year int, round int, offset int, limit int) (types.QualifyingData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.QUALIFYING)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return QualifyingRequest(path, query)
}
