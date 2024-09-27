package finishing_status

import (
	"fmt"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

// ergast.com/api/f1/status.json
func List(offset int, limit int) (types.FinishingStatusData, error) {
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.FinishingStatusData](urls.FINISHING_STATUS, query)
}

// ergast.com/api/f1/{year}/status.json
func BySeason(year int, offset int, limit int) (types.FinishingStatusData, error) {
	path := fmt.Sprintf("%d/%s", year, urls.FINISHING_STATUS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.FinishingStatusData](path, query)
}

// ergast.com/api/f1/{year}/{round}/status.json
func ByRace(year int, round int, offset int, limit int) (types.FinishingStatusData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.FINISHING_STATUS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.FinishingStatusData](path, query)
}
