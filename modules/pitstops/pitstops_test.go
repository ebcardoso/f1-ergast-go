package pitstops

import (
	"testing"

	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/stretchr/testify/assert"
)

func TestByRace(t *testing.T) {
	assert := assert.New(t)

	result, err := ByRace(2023, 11, 0, 50)
	assert.Nil(err, "Should not have errors")

	//Should not have an empty response
	assert.NotEqual(result, types.PitstopsData{})
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
		assert.NotEqual(race.Pitstops, "")
		for _, pits := range race.Pitstops {
			assert.NotEqual(pits.DriverId, "")
			assert.NotEqual(pits.Lap, "")
			assert.NotEqual(pits.Stop, "")
			assert.NotEqual(pits.Time, "")
			assert.NotEqual(pits.Duration, "")
		}
	}
}

func TestGetByNumber(t *testing.T) {
	assert := assert.New(t)

	result, err := GetByNumber(2023, 11, 2, 0, 50)
	assert.Nil(err, "Should not have errors")

	//Should not have an empty response
	assert.NotEqual(result, types.PitstopsData{})
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
	assert.NotEqual(result.RaceTable.Stop, "")
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
		assert.NotEqual(race.Time, "")
		assert.NotEqual(race.Pitstops, "")
		for _, pits := range race.Pitstops {
			assert.NotEqual(pits.DriverId, "")
			assert.NotEqual(pits.Lap, "")
			assert.NotEqual(pits.Stop, "")
			assert.NotEqual(pits.Time, "")
			assert.NotEqual(pits.Duration, "")
		}
	}
}
