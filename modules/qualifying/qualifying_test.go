package qualifying

import (
	"testing"

	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/stretchr/testify/assert"
)

func TestByRace(t *testing.T) {
	assert := assert.New(t)

	result, err := ByRace(2010, 11, 0, 30)
	assert.Nil(err, "Should not have errors")

	assert.NotEqual(result, types.QualifyingData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotEqual(result.RaceTable, "")
	assert.NotEqual(result.RaceTable.Season, "")
	assert.NotEqual(result.RaceTable.Round, "")
	assert.NotEqual(result.RaceTable.Races, "")
	for _, race := range result.RaceTable.Races {
		assert.NotEqual(race.Season, "")
		assert.NotEqual(race.Round, "")
		assert.NotEqual(race.Url, "")
		assert.NotEqual(race.RaceName, "")
		assert.NotEqual(race.Circuit, "")
		assert.NotEqual(race.Circuit.CircuitId, "")
		assert.NotEqual(race.Circuit.Url, "")
		assert.NotEqual(race.Circuit.CircuitName, "")
		assert.NotEqual(race.Circuit.Location.Lat, "")
		assert.NotEqual(race.Circuit.Location.Long, "")
		assert.NotEqual(race.Circuit.Location.Locality, "")
		assert.NotEqual(race.Circuit.Location.Country, "")
		assert.NotEqual(race.Date, "")
		assert.NotEqual(race.Time, "")
		assert.NotEqual(race.QualifyingResults, "")
		for _, qr := range race.QualifyingResults {
			assert.NotEqual(qr.Number, "")
			assert.NotEqual(qr.Position, "")
			assert.NotEqual(qr.Driver, "")
			assert.NotEqual(qr.Driver.DriverId, "")
			assert.NotEqual(qr.Driver.URL, "")
			assert.NotEqual(qr.Driver.GivenName, "")
			assert.NotEqual(qr.Driver.FamilyName, "")
			assert.NotEqual(qr.Driver.DateOfBirth, "")
			assert.NotEqual(qr.Driver.Nationality, "")
			assert.NotEqual(qr.Constructor.ConstructorId, "")
			assert.NotEqual(qr.Constructor.URL, "")
			assert.NotEqual(qr.Constructor.Name, "")
			assert.NotEqual(qr.Constructor.Nationality, "")
		}
	}
}
