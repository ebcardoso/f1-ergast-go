package results

import (
	"testing"

	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/stretchr/testify/assert"
)

func TestByRace(t *testing.T) {
	assert := assert.New(t)

	result, err := ByRace(1988, 02)
	assert.Nil(err, "Should not have errors")

	//Should not have an empty response
	assert.NotEqual(result, types.ResultsData{})

	//Validating fields
	assert.NotEqual(result, types.ConstructorsData{})
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
	for _, item := range result.RaceTable.Races {
		assert.NotEqual(item.Season, "")
		assert.NotEqual(item.Round, "")
		assert.NotEqual(item.Url, "")
		assert.NotEqual(item.RaceName, "")
		assert.NotEqual(item.Circuit, "")
		assert.NotEqual(item.Circuit.CircuitId, "")
		assert.NotEqual(item.Circuit.Url, "")
		assert.NotEqual(item.Circuit.CircuitName, "")
		assert.NotEqual(item.Circuit.Location, "")
		assert.NotEqual(item.Circuit.Location.Lat, "")
		assert.NotEqual(item.Circuit.Location.Long, "")
		assert.NotEqual(item.Circuit.Location.Locality, "")
		assert.NotEqual(item.Circuit.Location.Country, "")
		assert.NotEqual(item.Circuit.Location.Country, "")
		assert.NotEqual(item.Date, "")
		assert.NotEqual(item.Results, "")
		for _, res := range item.Results {
			assert.NotEqual(res.Number, "")
			assert.NotEqual(res.Position, "")
			assert.NotEqual(res.PositionText, "")
			assert.NotEqual(res.Points, "")
			assert.NotEqual(res.Driver, "")
			assert.NotEqual(res.Driver.DriverId, "")
			assert.NotEqual(res.Driver.URL, "")
			assert.NotEqual(res.Driver.GivenName, "")
			assert.NotEqual(res.Driver.FamilyName, "")
			assert.NotEqual(res.Driver.DateOfBirth, "")
			assert.NotEqual(res.Driver.Nationality, "")
			assert.NotEqual(res.Constructor, "")
			assert.NotEqual(res.Constructor.ConstructorId, "")
			assert.NotEqual(res.Constructor.URL, "")
			assert.NotEqual(res.Constructor.Name, "")
			assert.NotEqual(res.Constructor.Nationality, "")
			assert.NotEqual(res.Grid, "")
			assert.NotEqual(res.Laps, "")
			assert.NotEqual(res.Status, "")
		}
	}
}

func TestMostRecent(t *testing.T) {
	assert := assert.New(t)

	result, err := MostRecent()
	assert.Nil(err, "Should not have errors")

	//Should not have an empty response
	assert.NotEqual(result, types.ResultsData{})

	//Validating fields
	assert.NotEqual(result, types.ConstructorsData{})
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
	for _, item := range result.RaceTable.Races {
		assert.NotEqual(item.Season, "")
		assert.NotEqual(item.Round, "")
		assert.NotEqual(item.Url, "")
		assert.NotEqual(item.RaceName, "")
		assert.NotEqual(item.Circuit, "")
		assert.NotEqual(item.Circuit.CircuitId, "")
		assert.NotEqual(item.Circuit.Url, "")
		assert.NotEqual(item.Circuit.CircuitName, "")
		assert.NotEqual(item.Circuit.Location, "")
		assert.NotEqual(item.Circuit.Location.Lat, "")
		assert.NotEqual(item.Circuit.Location.Long, "")
		assert.NotEqual(item.Circuit.Location.Locality, "")
		assert.NotEqual(item.Circuit.Location.Country, "")
		assert.NotEqual(item.Circuit.Location.Country, "")
		assert.NotEqual(item.Date, "")
		assert.NotEqual(item.Time, "")
		assert.NotEqual(item.Results, "")
		for _, res := range item.Results {
			assert.NotEqual(res.Number, "")
			assert.NotEqual(res.Position, "")
			assert.NotEqual(res.PositionText, "")
			assert.NotEqual(res.Points, "")
			assert.NotEqual(res.Driver, "")
			assert.NotEqual(res.Driver.DriverId, "")
			assert.NotEqual(res.Driver.PermanentNumber, "")
			assert.NotEqual(res.Driver.URL, "")
			assert.NotEqual(res.Driver.GivenName, "")
			assert.NotEqual(res.Driver.FamilyName, "")
			assert.NotEqual(res.Driver.DateOfBirth, "")
			assert.NotEqual(res.Driver.Nationality, "")
			assert.NotEqual(res.Constructor, "")
			assert.NotEqual(res.Constructor.ConstructorId, "")
			assert.NotEqual(res.Constructor.URL, "")
			assert.NotEqual(res.Constructor.Name, "")
			assert.NotEqual(res.Constructor.Nationality, "")
			assert.NotEqual(res.Grid, "")
			assert.NotEqual(res.Laps, "")
			assert.NotEqual(res.Status, "")
		}
	}
}
