package finishing_status

import (
	"testing"

	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	assert := assert.New(t)

	result, err := List(0, 50)
	assert.Nil(err, "Should not have errors")

	//Should not have an empty response
	assert.NotEqual(result, types.FinishStatus{}, "Should not have an empty response")

	assert.NotEqual(result, types.ConstructorsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotEqual(result.StatusTable, "")
	assert.NotEqual(result.StatusTable.Status, "")
	for _, status := range result.StatusTable.Status {
		assert.NotEqual(status.StatusId, "")
		assert.NotEqual(status.Count, "")
		assert.NotEqual(status.Status, "")
	}
}

func TestBySeason(t *testing.T) {
	assert := assert.New(t)

	result, err := BySeason(2023, 0, 50)
	assert.Nil(err, "Should not have errors")

	//Should not have an empty response
	assert.NotEqual(result, types.FinishStatus{}, "Should not have an empty response")

	assert.NotEqual(result, types.ConstructorsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotEqual(result.StatusTable, "")
	assert.NotEqual(result.StatusTable.Season, "")
	assert.NotEqual(result.StatusTable.Status, "")
	for _, status := range result.StatusTable.Status {
		assert.NotEqual(status.StatusId, "")
		assert.NotEqual(status.Count, "")
		assert.NotEqual(status.Status, "")
	}
}

func TestByRace(t *testing.T) {
	assert := assert.New(t)

	result, err := ByRace(2022, 12, 0, 50)
	assert.Nil(err, "Should not have errors")

	//Should not have an empty response
	assert.NotEqual(result, types.FinishStatus{}, "Should not have an empty response")

	assert.NotEqual(result, types.ConstructorsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotEqual(result.StatusTable, "")
	assert.NotEqual(result.StatusTable.Season, "")
	assert.NotEqual(result.StatusTable.Round, "")
	assert.NotEqual(result.StatusTable.Status, "")
	for _, status := range result.StatusTable.Status {
		assert.NotEqual(status.StatusId, "")
		assert.NotEqual(status.Count, "")
		assert.NotEqual(status.Status, "")
	}
}
