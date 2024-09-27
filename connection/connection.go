package connection

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

func SendRequestGet[T any](path string, query string) (T, error) {
	var response types.ErgastResponse[T]

	url := fmt.Sprintf("%s%s.json%s", urls.ROOT_URL, path, query)

	// Sending request
	resp, err := http.Get(url)
	if err != nil {
		return response.Data, err
	}

	// Reading response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()

	// Parsing response to json
	err = json.Unmarshal(body, &response)
	if err != nil {
		return response.Data, err
	}
	return response.Data, nil
}
