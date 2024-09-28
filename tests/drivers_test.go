package tests

import (
	"testing"

	"github.com/ebcardoso/f1-ergast-go/modules/f1_services"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/stretchr/testify/assert"
)

func TestListDrivers(t *testing.T) {
	assert := assert.New(t)

	result, err := f1_services.ListDrivers(0, 50)
	assert.Nil(err, "Should not be errors")

	assert.NotEqual(result, types.ConstructorsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotNil(result.DriverTable)
	assert.NotNil(result.DriverTable.Drivers)
	for _, item := range result.DriverTable.Drivers {
		assert.NotEqual(item.DriverId, "")
		assert.NotEqual(item.URL, "")
		assert.NotEqual(item.GivenName, "")
		assert.NotEqual(item.FamilyName, "")
		assert.NotEqual(item.DateOfBirth, "")
		assert.NotEqual(item.Nationality, "")
	}
}

func TestListDriversBySeason(t *testing.T) {
	assert := assert.New(t)

	result, err := f1_services.ListDriversBySeason(2023, 0, 30)
	assert.Nil(err, "Should not be errors")

	assert.NotEqual(result, types.ConstructorsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotNil(result.DriverTable)
	assert.NotEqual(result.DriverTable.Season, "")
	assert.NotNil(result.DriverTable.Drivers)
	for _, item := range result.DriverTable.Drivers {
		assert.NotEqual(item.DriverId, "")
		assert.NotEqual(item.URL, "")
		assert.NotEqual(item.GivenName, "")
		assert.NotEqual(item.FamilyName, "")
		assert.NotEqual(item.DateOfBirth, "")
		assert.NotEqual(item.Nationality, "")
	}
}

func TestListDriversByRace(t *testing.T) {
	assert := assert.New(t)

	result, err := f1_services.ListDriversByRace(2023, 12, 0, 30)
	assert.Nil(err, "Should not be errors")

	assert.NotEqual(result, types.ConstructorsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotNil(result.DriverTable)
	assert.NotEqual(result.DriverTable.Season, "")
	assert.NotEqual(result.DriverTable.Round, "")
	assert.NotNil(result.DriverTable.Drivers)
	for _, item := range result.DriverTable.Drivers {
		assert.NotEqual(item.DriverId, "")
		assert.NotEqual(item.URL, "")
		assert.NotEqual(item.GivenName, "")
		assert.NotEqual(item.FamilyName, "")
		assert.NotEqual(item.DateOfBirth, "")
		assert.NotEqual(item.Nationality, "")
	}
}

func TestFindDriverById(t *testing.T) {
	assert := assert.New(t)

	result, err := f1_services.FindDriverById("raikkonen")
	assert.Nil(err, "Should not be errors")

	assert.NotEqual(result, types.ConstructorsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotNil(result.DriverTable)
	assert.NotEqual(result.DriverTable.DriverId, "")
	assert.NotNil(result.DriverTable.Drivers)
	for _, item := range result.DriverTable.Drivers {
		assert.NotEqual(item.DriverId, "")
		assert.NotEqual(item.URL, "")
		assert.NotEqual(item.GivenName, "")
		assert.NotEqual(item.FamilyName, "")
		assert.NotEqual(item.DateOfBirth, "")
		assert.NotEqual(item.Nationality, "")
	}
}
