package f1_services

import (
	"fmt"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

// ergast.com/api/f1/current/last/results.json
func MostRecentResults() (types.ResultsData, error) {
	return connection.SendRequestGet[types.ResultsData](urls.RESULT_MOST_RECENT, "")
}

// ergast.com/api/f1/{year}/{round}/results.json
func ListResultsByRace(year int, round int) (types.ResultsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.RESULT_RACE)
	return connection.SendRequestGet[types.ResultsData](path, "")
}
