package f1_services

import (
	"fmt"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

/* DRIVERS */

// ergast.com/api/f1/current/driverStandings.json
func DriversCurrentStandingss() (types.StandingsData, error) {
	path := fmt.Sprintf("%s/%s", urls.STANDINGS_CURRENT, urls.STANDINGS_DRIVERS)

	return connection.SendRequestGet[types.StandingsData](path, "")
}

// ergast.com/api/f1/{{year}}/driverStandings.json
func DriversStandingsBySeason(year int, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%d/%s", year, urls.STANDINGS_DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}

// ergast.com/api/f1/{{year}}/{{round}}/driverStandings.json
func DriversStandingsAfterRace(year int, round int, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.STANDINGS_DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}

// ergast.com/api/f1/drivers/{{driverid}}/driverStandings.json
func StandingsHistoryByDriver(driverId string, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%s/%s/%s", urls.DRIVERS, driverId, urls.STANDINGS_DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}

// ergast.com/api/f1/driverStandings/1.json
func ChampionsHistoryByDrivers(offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%s/1", urls.STANDINGS_DRIVERS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}

/* CONSTRUCTORS */

// ergast.com/api/f1/current/constructorStandings.json
func ConstructorsCurrentStandings() (types.StandingsData, error) {
	path := fmt.Sprintf("%s/%s", urls.STANDINGS_CURRENT, urls.STANDINGS_CONSTRUCTORS)

	return connection.SendRequestGet[types.StandingsData](path, "")
}

// ergast.com/api/f1/{{year}}/constructorStandings.json
func ConstructorsStandingsBySeason(year int, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%d/%s", year, urls.STANDINGS_CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}

// ergast.com/api/f1/{{year}}/{{round}}/constructorStandings.json
func ConstructorsStandingsAfterRace(year int, round int, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.STANDINGS_CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}

// ergast.com/api/f1/constructors/{{constructorid}}/constructorStandings.json
func StandingsHistoryByConstructor(constructorId string, offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%s/%s/%s", urls.CONSTRUCTORS, constructorId, urls.STANDINGS_CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}

// ergast.com/api/f1/constructorStandings/1.json
func ChampionsHistoryByConstructors(offset int, limit int) (types.StandingsData, error) {
	path := fmt.Sprintf("%s/1", urls.STANDINGS_CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.StandingsData](path, query)
}
