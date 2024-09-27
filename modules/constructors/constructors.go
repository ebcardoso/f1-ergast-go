package constructors

import (
	"fmt"

	"github.com/ebcardoso/f1-ergast-go/connection"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/ebcardoso/f1-ergast-go/utils/urls"
)

// ergast.com/api/f1/constructors.json
func List(offset int, limit int) (types.ConstructorsData, error) {
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.ConstructorsData](urls.CONSTRUCTORS, query)
}

// ergast.com/api/f1/{year}/constructors.json
func BySeason(year int, offset int, limit int) (types.ConstructorsData, error) {
	path := fmt.Sprintf("%d/%s", year, urls.CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.ConstructorsData](path, query)
}

// ergast.com/api/f1/{year}/{round}/constructors.json
func ByRace(year int, round int, offset int, limit int) (types.ConstructorsData, error) {
	path := fmt.Sprintf("%d/%d/%s", year, round, urls.CONSTRUCTORS)
	query := fmt.Sprintf("%s%d%s%d", "?offset=", offset, "&limit=", limit)

	return connection.SendRequestGet[types.ConstructorsData](path, query)
}

// ergast.com/api/f1/{constructorsId}/constructors.json
func GetByConstructorId(constructorId string) (types.ConstructorsData, error) {
	path := fmt.Sprintf("%s/%s", urls.CONSTRUCTORS, constructorId)

	return connection.SendRequestGet[types.ConstructorsData](path, "")
}
