package qualifying

import (
	"fmt"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

// ergast.com/api/f1/{year}/{round}/qualifying.json
func ByRace(year int, round int, offset int, limit int) (types.QualifyingData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.QUALIFYING)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.QualifyingData](path, query)
}
