package circuits

import (
	"testing"

	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	assert := assert.New(t)

	result, err := List(0, 50)
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

func TestBySeason(t *testing.T) {
	assert := assert.New(t)

	result, err := BySeason(2023, 0, 30)
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

func TestByRace(t *testing.T) {
	assert := assert.New(t)

	result, err := ByRace(2023, 12, 0, 30)
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

func TestGetByCircuitId(t *testing.T) {
	assert := assert.New(t)

	result, err := GetByCircuitId("adelaide")
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
