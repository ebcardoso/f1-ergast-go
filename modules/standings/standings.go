package standings

import (
	"fmt"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

// ergast.com/api/f1/{{year}}/{{round}}/driverStandings.json
func DriversAfterRace(year int, round int, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.STANDINGS_DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}

// ergast.com/api/f1/{{year}}/{{round}}/constructorStandings.json
func ConstructorsAfterRace(year int, round int, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.STANDINGS_CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}

// ergast.com/api/f1/{{year}}/driverStandings.json
func DriversBySeason(year int, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%d/%s", year, urls.STANDINGS_DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}

// ergast.com/api/f1/{{year}}/constructorStandings.json
func ConstructorsBySeason(year int, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%d/%s", year, urls.STANDINGS_CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}

// ergast.com/api/f1/current/driverStandings.json
func DriversCurrent() (types.StandingsData, error) {
	path := fmt.Sprintf("%s/%s", urls.STANDINGS_CURRENT, urls.STANDINGS_DRIVERS)

	return connection.SendRequestGet[types.StandingsData](path, "")
}

// ergast.com/api/f1/current/constructorStandings.json
func ConstructorsCurrent() (types.StandingsData, error) {
	path := fmt.Sprintf("%s/%s", urls.STANDINGS_CURRENT, urls.STANDINGS_CONSTRUCTORS)

	return connection.SendRequestGet[types.StandingsData](path, "")
}

// ergast.com/api/f1/drivers/{{driverid}}/driverStandings.json
func HistoryByDriver(driverId string, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%s/%s/%s", urls.DRIVERS, driverId, urls.STANDINGS_DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}

// ergast.com/api/f1/constructors/{{constructorid}}/constructorStandings.json
func HistoryByConstructor(constructorId string, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%s/%s/%s", urls.CONSTRUCTORS, constructorId, urls.STANDINGS_CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}

// ergast.com/api/f1/driverStandings/1.json
func WinnersByDrivers(offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%s/1", urls.STANDINGS_DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}

// ergast.com/api/f1/constructorStandings/1.json
func WinnersByConstructors(offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%s/1", urls.STANDINGS_CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}
