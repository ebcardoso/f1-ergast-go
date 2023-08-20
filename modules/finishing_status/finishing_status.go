package finishing_status

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils"
)

func FinishingStatusRequest(path string, query string) (types.FinishingStatusData, error) {
	resp, err := connection.Get(path, query)
	if err != nil {
		return types.FinishingStatusData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.FinishingStatusData{}, err
	}
	defer resp.Body.Close()

	var response types.FinishingStatusResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return types.FinishingStatusData{}, err
	}
	return response.Data, nil
}

// ergast.com/api/f1/status.json
func List(offset int, limit int) (types.FinishingStatusData, error) {
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return FinishingStatusRequest(utils.FINISHING_STATUS, query)
}

// ergast.com/api/f1/{year}/status.json
func BySeason(year int, offset int, limit int) (types.FinishingStatusData, error) {
	path := fmt.Sprintf("%d/%s", year, utils.FINISHING_STATUS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return FinishingStatusRequest(path, query)
}

// ergast.com/api/f1/{year}/{round}/status.json
func ByRace(year int, round int, offset int, limit int) (types.FinishingStatusData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, utils.FINISHING_STATUS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return FinishingStatusRequest(path, query)
}
