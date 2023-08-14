package seasons

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils"
)

func SeasonRequest(path string, query string) (types.SeasonsData, error) {
	resp, err := connection.Get(path, query)
	if err != nil {
		return types.SeasonsData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.SeasonsData{}, err
	}
	defer resp.Body.Close()

	var response types.SeasonsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return types.SeasonsData{}, err
	}
	return response.Data, nil
}

func List(offset int, limit int) (types.SeasonsData, error) {
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return SeasonRequest(utils.SEASONS, query)
}
