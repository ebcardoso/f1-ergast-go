package standings

import (
	"testing"

	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/stretchr/testify/assert"
)

func TestDriversAfterRace(t *testing.T) {
	assert := assert.New(t)

	result, err := DriversAfterRace(2023, 1, 0, 25)
	assert.Nil(err, "Should not have errors")

	assert.NotEqual(result, types.StandingsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotEqual(result.StandingsTable, "")
	assert.NotEqual(result.StandingsTable.Season, "")
	assert.NotEqual(result.StandingsTable.Round, "")
	assert.NotEqual(result.StandingsTable.StandingsLists, "")
	for _, sl := range result.StandingsTable.StandingsLists {
		assert.NotEqual(sl.Season, "")
		assert.NotEqual(sl.Round, "")
		assert.NotEqual(sl.DriverStandings, "")
		for _, ds := range sl.DriverStandings {
			assert.NotEqual(ds.Position, "")
			assert.NotEqual(ds.PositionText, "")
			assert.NotEqual(ds.Points, "")
			assert.NotEqual(ds.Wins, "")
			assert.NotEqual(ds.Driver, "")
			assert.NotEqual(ds.Driver.DriverId, "")
			assert.NotEqual(ds.Driver.URL, "")
			assert.NotEqual(ds.Driver.GivenName, "")
			assert.NotEqual(ds.Driver.FamilyName, "")
			assert.NotEqual(ds.Driver.DateOfBirth, "")
			assert.NotEqual(ds.Driver.Nationality, "")
			assert.NotEqual(ds.Constructors, "")
			for _, c := range ds.Constructors {
				assert.NotEqual(c.ConstructorId, "")
				assert.NotEqual(c.URL, "")
				assert.NotEqual(c.Name, "")
				assert.NotEqual(c.Nationality, "")
			}
		}
	}
}

func TestConstructorsAfterRace(t *testing.T) {
	assert := assert.New(t)

	result, err := ConstructorsAfterRace(2023, 1, 0, 15)
	assert.Nil(err, "Should not have errors")

	assert.NotEqual(result, types.StandingsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotEqual(result.StandingsTable, "")
	assert.NotEqual(result.StandingsTable.Season, "")
	assert.NotEqual(result.StandingsTable.Round, "")
	assert.NotEqual(result.StandingsTable.StandingsLists, "")
	for _, sl := range result.StandingsTable.StandingsLists {
		assert.NotEqual(sl.Season, "")
		assert.NotEqual(sl.Round, "")
		assert.NotEqual(sl.ConstructorStandings, "")
		for _, cs := range sl.ConstructorStandings {
			assert.NotEqual(cs.Position, "")
			assert.NotEqual(cs.PositionText, "")
			assert.NotEqual(cs.Points, "")
			assert.NotEqual(cs.Wins, "")
			assert.NotEqual(cs.Constructor, "")
			assert.NotEqual(cs.Constructor.ConstructorId, "")
			assert.NotEqual(cs.Constructor.URL, "")
			assert.NotEqual(cs.Constructor.Name, "")
			assert.NotEqual(cs.Constructor.Nationality, "")
		}
	}
}
