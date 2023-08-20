package laps

import (
	"testing"

	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGetLapTimes(t *testing.T) {
	assert := assert.New(t)

	result, err := GetLapTimes(2020, 11, 1, 0, 50)
	assert.Nil(err, "Should not have errors")

	assert.NotEqual(result, types.LapsData{})
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
		assert.NotEqual(race.Laps, "")
		for _, lap := range race.Laps {
			assert.NotEqual(lap.Number, "")
			assert.NotEqual(lap.Timings, "")
			for _, timing := range lap.Timings {
				assert.NotEqual(timing.DriverId, "")
				assert.NotEqual(timing.Position, "")
				assert.NotEqual(timing.Time, "")
			}
		}
	}
}
