package schedules

import (
	"fmt"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

// ergast.com/api/f1/{year}.json
func BySeason(year int, offset int, limit int) (types.RacesData, error) {
	path := fmt.Sprintf("%d", year)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.RacesData](path, query)
}

// ergast.com/api/f1/current.json
func ByCurrentSeason(offset int, limit int) (types.RacesData, error) {
	path := urls.SCHEDULE_CURRENT
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.RacesData](path, query)
}

// ergast.com/api/f1/{year}/{round}.json
func ByRace(year int, round int) (types.RacesData, error) {
	path := fmt.Sprintf("%d/%d", year, round)
	query := ""

	return connection.SendRequestGet[types.RacesData](path, query)
}
