package constructors

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
	assert.NotNil(result.ConstructorTable)
	assert.NotNil(result.ConstructorTable.Constructors)
	for _, item := range result.ConstructorTable.Constructors {
		assert.NotEqual(item.ConstructorId, "")
		assert.NotEqual(item.URL, "")
		assert.NotEqual(item.Name, "")
		assert.NotEqual(item.Nationality, "")
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
	assert.NotNil(result.ConstructorTable)
	assert.NotEqual(result.ConstructorTable.Season, "")
	assert.NotNil(result.ConstructorTable.Constructors)
	for _, item := range result.ConstructorTable.Constructors {
		assert.NotEqual(item.ConstructorId, "")
		assert.NotEqual(item.URL, "")
		assert.NotEqual(item.Name, "")
		assert.NotEqual(item.Nationality, "")
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
	assert.NotNil(result.ConstructorTable)
	assert.NotEqual(result.ConstructorTable.Season, "")
	assert.NotEqual(result.ConstructorTable.Round, "")
	assert.NotNil(result.ConstructorTable.Constructors)
	for _, item := range result.ConstructorTable.Constructors {
		assert.NotEqual(item.ConstructorId, "")
		assert.NotEqual(item.URL, "")
		assert.NotEqual(item.Name, "")
		assert.NotEqual(item.Nationality, "")
	}
}

func TestGetByConstructorId(t *testing.T) {
	assert := assert.New(t)

	result, err := GetByConstructorId("mclaren")
	assert.Nil(err, "Should not be errors")

	assert.NotEqual(result, types.ConstructorsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotNil(result.ConstructorTable)
	assert.NotEqual(result.ConstructorTable.ConstructorId, "")
	assert.NotNil(result.ConstructorTable.Constructors)
	for _, item := range result.ConstructorTable.Constructors {
		assert.NotEqual(item.ConstructorId, "")
		assert.NotEqual(item.URL, "")
		assert.NotEqual(item.Name, "")
		assert.NotEqual(item.Nationality, "")
	}
}
