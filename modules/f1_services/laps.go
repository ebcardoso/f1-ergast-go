package f1_services

import (
	"fmt"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

// ergast.com/api/f1/{year}/{round}/laps/{lap}.json
func ListLapTimes(year int, round int, lap int, offset int, limit int) (types.LapsData, error) {
	path := fmt.Sprintf("%d/%d/%s/%d", year, round, urls.LAPS, lap)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.LapsData](path, query)
}
