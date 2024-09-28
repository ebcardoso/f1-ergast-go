package tests

import (
	"testing"

	"github.com/ebcardoso/f1-ergast-go/modules/f1_services"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/stretchr/testify/assert"
)

func TestListSchedulesBySeason(t *testing.T) {
	assert := assert.New(t)

	result, err := f1_services.ListSchedulesBySeason(2023, 0, 10)
	assert.Nil(err, "Should not have errors")

	assert.NotEqual(result, types.RacesData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotNil(result.RaceTable)
	assert.NotEqual(result.RaceTable.Season, "")
	assert.NotNil(result.RaceTable.Races)
	for _, item := range result.RaceTable.Races {
		assert.NotEqual(item.Season, "")
		assert.NotEqual(item.Round, "")
		assert.NotEqual(item.Url, "")
		assert.NotEqual(item.RaceName, "")
		assert.NotNil(item.Circuit, "")
		assert.NotEqual(item.Circuit.CircuitId, "")
		assert.NotEqual(item.Circuit.Url, "")
		assert.NotEqual(item.Circuit.CircuitName, "")
		assert.NotNil(item.Circuit.Location)
		assert.NotEqual(item.Circuit.Location.Lat, "")
		assert.NotEqual(item.Circuit.Location.Long, "")
		assert.NotEqual(item.Circuit.Location.Locality, "")
		assert.NotEqual(item.Circuit.Location.Country, "")
		assert.NotEqual(item.Date, "")
	}
}

func TestListSchedulesByCurrentSeason(t *testing.T) {
	assert := assert.New(t)

	result, err := f1_services.ListSchedulesByCurrentSeason(0, 10)
	assert.Nil(err, "Should not have errors")

	assert.NotEqual(result, types.RacesData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotNil(result.RaceTable)
	assert.NotEqual(result.RaceTable.Season, "")
	assert.NotNil(result.RaceTable.Races)
	for _, item := range result.RaceTable.Races {
		assert.NotEqual(item.Season, "")
		assert.NotEqual(item.Round, "")
		assert.NotEqual(item.Url, "")
		assert.NotEqual(item.RaceName, "")
		assert.NotNil(item.Circuit, "")
		assert.NotEqual(item.Circuit.CircuitId, "")
		assert.NotEqual(item.Circuit.Url, "")
		assert.NotEqual(item.Circuit.CircuitName, "")
		assert.NotNil(item.Circuit.Location)
		assert.NotEqual(item.Circuit.Location.Lat, "")
		assert.NotEqual(item.Circuit.Location.Long, "")
		assert.NotEqual(item.Circuit.Location.Locality, "")
		assert.NotEqual(item.Circuit.Location.Country, "")
		assert.NotEqual(item.Date, "")
	}
}

func TestListSchedulesByRace(t *testing.T) {
	assert := assert.New(t)

	result, err := f1_services.ListSchedulesByRace(2023, 1)
	assert.Nil(err, "Should not have errors")

	assert.NotEqual(result, types.RacesData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotNil(result.RaceTable)
	assert.NotEqual(result.RaceTable.Season, "")
	assert.NotEqual(result.RaceTable.Round, "")
	assert.NotNil(result.RaceTable.Races)
	for _, item := range result.RaceTable.Races {
		assert.NotEqual(item.Season, "")
		assert.NotEqual(item.Round, "")
		assert.NotEqual(item.Url, "")
		assert.NotEqual(item.RaceName, "")
		assert.NotNil(item.Circuit, "")
		assert.NotEqual(item.Circuit.CircuitId, "")
		assert.NotEqual(item.Circuit.Url, "")
		assert.NotEqual(item.Circuit.CircuitName, "")
		assert.NotNil(item.Circuit.Location)
		assert.NotEqual(item.Circuit.Location.Lat, "")
		assert.NotEqual(item.Circuit.Location.Long, "")
		assert.NotEqual(item.Circuit.Location.Locality, "")
		assert.NotEqual(item.Circuit.Location.Country, "")
		assert.NotEqual(item.Date, "")
	}
}
