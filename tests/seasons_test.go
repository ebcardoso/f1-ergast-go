package tests

import (
	"testing"

	"github.com/ebcardoso/f1-ergast-go/modules/f1_services"
	"github.com/ebcardoso/f1-ergast-go/types"
	"github.com/stretchr/testify/assert"
)

func TestListSeasons(t *testing.T) {
	assert := assert.New(t)

	result, err := f1_services.ListSeasons(0, 100)
	assert.Nil(err, "Should not have errors")

	assert.NotEqual(result, types.SeasonsData{})
	assert.NotEqual(result.Limit, "")
	assert.NotEqual(result.Offset, "")
	assert.NotEqual(result.Total, "")
	assert.NotEqual(result.Series, "")
	assert.NotEqual(result.Url, "")
	assert.NotEqual(result.Xmlns, "")
	assert.NotNil(result.SeasonTable)
	assert.NotNil(result.SeasonTable.Seasons)
	for _, item := range result.SeasonTable.Seasons {
		assert.NotEqual(item.Season, "")
		assert.NotEqual(item.URL, "")
	}
}
