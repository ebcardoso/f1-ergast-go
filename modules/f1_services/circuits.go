package f1_services

import (
	"fmt"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

// ergast.com/api/f1/circuits.json
func ListCircuits(offset int, limit int) (types.CircuitsData, error) {
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.CircuitsData](urls.CIRCUITS, query)
}

// ergast.com/api/f1/{year}/circuits.json
func ListCircuitsBySeason(year int, offset int, limit int) (types.CircuitsData, error) {
	path := fmt.Sprintf("%d/%s", year, urls.CIRCUITS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.CircuitsData](path, query)
}

// ergast.com/api/f1/{year}/{round}/circuits.json
func ListCircuitsByRace(year int, round int, offset int, limit int) (types.CircuitsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.CIRCUITS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.CircuitsData](path, query)
}

// ergast.com/api/f1/{circuitId}/circuits.json
func FindCircuitById(circuitId string) (types.CircuitsData, error) {
	path := fmt.Sprintf("%s/%s", urls.CIRCUITS, circuitId)

	return connection.SendRequestGet[types.CircuitsData](path, "")
}
