package results

import (
	"fmt"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

// ergast.com/api/f1/{year}/{round}/results.json
func ByRace(year int, round int) (types.ResultsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.RESULT_RACE)
	return connection.SendRequestGet[types.ResultsData](path, "")
}

// ergast.com/api/f1/current/last/results.json
func MostRecent() (types.ResultsData, error) {
	return connection.SendRequestGet[types.ResultsData](urls.RESULT_MOST_RECENT, "")
}
