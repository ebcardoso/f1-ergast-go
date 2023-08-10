package constructors

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/utils"
)

type ConstructorsResponse struct {
	MRData struct {
		Limit            string `json:"limit,omitempty"`
		Offset           string `json:"offset,omitempty"`
		Total            string `json:"total,omitempty"`
		Series           string `json:"series,omitempty"`
		Url              string `json:"url,omitempty"`
		Xmlns            string `json:"xmlns,omitempty"`
		ConstructorTable struct {
			Season        string `json:"season,omitempty"`
			Round         string `json:"round,omitempty"`
			ConstructorId string `json:"constructorId,omitempty"`
			Constructors  []struct {
				ConstructorId string `json:"constructorId,omitempty"`
				URL           string `json:"url,omitempty"`
				Name          string `json:"name,omitempty"`
				Nationality   string `json:"nationality,omitempty"`
			}
		} `json:"ConstructorTable,omitempty"`
	} `json:"MRData,omitempty"`
}

func ConstructorRequest(path string, query string) (ConstructorsResponse, error) {
	resp, err := connection.Get(path, query)
	if err != nil {
		return ConstructorsResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ConstructorsResponse{}, err
	}
	defer resp.Body.Close()

	var response ConstructorsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ConstructorsResponse{}, err
	}
	return response, nil
}

// ergast.com/api/f1/constructors.json
func List(offset int, limit int) (ConstructorsResponse, error) {
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return ConstructorRequest(utils.CONSTRUCTORS, query)
}

// ergast.com/api/f1/{year}/constructors.json
func BySeason(year int, offset int, limit int) (ConstructorsResponse, error) {
	path := fmt.Sprintf("%d/%s", year, utils.CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return ConstructorRequest(path, query)
}

// ergast.com/api/f1/{year}/{round}/constructors.json
func ByRace(year int, round int, offset int, limit int) (ConstructorsResponse, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, utils.CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return ConstructorRequest(path, query)
}

// ergast.com/api/f1/{constructorsId}/constructors.json
func GetByConstructorId(constructorId string) (ConstructorsResponse, error) {
	path := fmt.Sprintf("%s/%s", utils.CONSTRUCTORS, constructorId)

	return ConstructorRequest(path, "")
}
