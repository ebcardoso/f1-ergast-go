package tests

import (
	"testing"

	"github.com/ebcardoso/f1-ergast-go/modules/f1_services"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/stretchr/testify/assert"
)

func TestListCircuits(t *testing.T) {
	assert := assert.New(t)

	result, err := f1_services.ListCircuits(0, 50)
	assert.Nil(err, "Should not be errors")

	assert.NotEqual(result, types.ConstructorsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotNil(result.CircuitTable)
	assert.NotNil(result.CircuitTable.Circuits)
	for _, item := range result.CircuitTable.Circuits {
		assert.NotEqual(item.CircuitId, "")
		assert.NotEqual(item.Url, "")
		assert.NotEqual(item.CircuitName, "")
		assert.NotEqual(item.Location.Lat, "")
		assert.NotEqual(item.Location.Long, "")
		assert.NotEqual(item.Location.Locality, "")
		assert.NotEqual(item.Location.Country, "")
	}
}

func TestListCircuitsBySeason(t *testing.T) {
	assert := assert.New(t)

	result, err := f1_services.ListCircuitsBySeason(2023, 0, 30)
	assert.Nil(err, "Should not be errors")

	assert.NotEqual(result, types.ConstructorsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotNil(result.CircuitTable)
	assert.NotNil(result.CircuitTable.Season)
	assert.NotNil(result.CircuitTable.Circuits)
	for _, item := range result.CircuitTable.Circuits {
		assert.NotEqual(item.CircuitId, "")
		assert.NotEqual(item.Url, "")
		assert.NotEqual(item.CircuitName, "")
		assert.NotEqual(item.Location.Lat, "")
		assert.NotEqual(item.Location.Long, "")
		assert.NotEqual(item.Location.Locality, "")
		assert.NotEqual(item.Location.Country, "")
	}
}

func TestListCircuitsByRace(t *testing.T) {
	assert := assert.New(t)

	result, err := f1_services.ListCircuitsByRace(2023, 12, 0, 30)
	assert.Nil(err, "Should not be errors")

	assert.NotEqual(result, types.ConstructorsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotNil(result.CircuitTable)
	assert.NotNil(result.CircuitTable.Season)
	assert.NotNil(result.CircuitTable.Round)
	assert.NotNil(result.CircuitTable.Circuits)
	for _, item := range result.CircuitTable.Circuits {
		assert.NotEqual(item.CircuitId, "")
		assert.NotEqual(item.Url, "")
		assert.NotEqual(item.CircuitName, "")
		assert.NotEqual(item.Location.Lat, "")
		assert.NotEqual(item.Location.Long, "")
		assert.NotEqual(item.Location.Locality, "")
		assert.NotEqual(item.Location.Country, "")
	}
}

func TestFindCircuitById(t *testing.T) {
	assert := assert.New(t)

	result, err := f1_services.FindCircuitById("adelaide")
	assert.Nil(err, "Should not be errors")

	assert.NotEqual(result, types.ConstructorsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotNil(result.CircuitTable)
	assert.NotNil(result.CircuitTable.CircuitId)
	assert.NotNil(result.CircuitTable.Circuits)
	for _, item := range result.CircuitTable.Circuits {
		assert.NotEqual(item.CircuitId, "")
		assert.NotEqual(item.Url, "")
		assert.NotEqual(item.CircuitName, "")
		assert.NotEqual(item.Location.Lat, "")
		assert.NotEqual(item.Location.Long, "")
		assert.NotEqual(item.Location.Locality, "")
		assert.NotEqual(item.Location.Country, "")
	}
}
