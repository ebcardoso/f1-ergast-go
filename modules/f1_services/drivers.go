package f1_services

import (
	"fmt"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

// ergast.com/api/f1/drivers.json
func ListDrivers(offset int, limit int) (types.DriversData, error) {
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.DriversData](urls.DRIVERS, query)
}

// ergast.com/api/f1/{year}/drivers.json
func ListDriversBySeason(year int, offset int, limit int) (types.DriversData, error) {
	path := fmt.Sprintf("%d/%s", year, urls.DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.DriversData](path, query)
}

// ergast.com/api/f1/{year}/{round}/drivers.json
func ListDriversByRace(year int, round int, offset int, limit int) (types.DriversData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.DriversData](path, query)
}

// ergast.com/api/f1/{driverId}/drivers.json
func FindDriverById(driverId string) (types.DriversData, error) {
	path := fmt.Sprintf("%s/%s", urls.DRIVERS, driverId)

	return connection.SendRequestGet[types.DriversData](path, "")
}
