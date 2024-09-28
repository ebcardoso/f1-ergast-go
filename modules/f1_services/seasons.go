package f1_services

import (
	"fmt"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

func ListSeasons(offset int, limit int) (types.SeasonsData, error) {
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.SeasonsData](urls.SEASONS, query)
}
