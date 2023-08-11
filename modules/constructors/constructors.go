package constructors

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils"
)

func ConstructorRequest(path string, query string) (types.ConstructorsData, error) {
	resp, err := connection.Get(path, query)
	if err != nil {
		return types.ConstructorsData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.ConstructorsData{}, err
	}
	defer resp.Body.Close()

	var response types.ConstructorsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return types.ConstructorsData{}, err
	}
	return response.Data, nil
}

// ergast.com/api/f1/constructors.json
func List(offset int, limit int) (types.ConstructorsData, error) {
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return ConstructorRequest(utils.CONSTRUCTORS, query)
}

// ergast.com/api/f1/{year}/constructors.json
func BySeason(year int, offset int, limit int) (types.ConstructorsData, error) {
	path := fmt.Sprintf("%d/%s", year, utils.CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return ConstructorRequest(path, query)
}

// ergast.com/api/f1/{year}/{round}/constructors.json
func ByRace(year int, round int, offset int, limit int) (types.ConstructorsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, utils.CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return ConstructorRequest(path, query)
}

// ergast.com/api/f1/{constructorsId}/constructors.json
func GetByConstructorId(constructorId string) (types.ConstructorsData, error) {
	path := fmt.Sprintf("%s/%s", utils.CONSTRUCTORS, constructorId)

	return ConstructorRequest(path, "")
}
