package pitstops

import (
	"fmt"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

// http://ergast.com/api/f1/{year}/{round}/pitstops.json
func ByRace(year int, round int, offset int, limit int) (types.PitstopsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.PITSTOPS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.PitstopsData](path, query)
}

// http://ergast.com/api/f1/{year}/{round}/pitstops/{number}.json
func GetByNumber(year int, round int, number int, offset int, limit int) (types.PitstopsData, error) {
	path := fmt.Sprintf("%d/%d/%s/%d", year, round, urls.PITSTOPS, number)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.PitstopsData](path, query)
}
