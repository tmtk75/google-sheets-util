package sheetsutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchCells(t *testing.T) {
	t.Run("capital!A1:B", func(t *testing.T) {
		cells, err := FetchCells(srv.Service, spreadsheetId, "capital!A1:B")

		assert.Nil(t, err)
		assert.Equal(t, 4, len(cells))
		assert.Equal(t, "country", cells[0][0].FormattedValue)
		assert.Equal(t, "capital", cells[0][1].FormattedValue)
		assert.Equal(t, "Korea", cells[3][0].FormattedValue)
		assert.Equal(t, "Soul", cells[3][1].FormattedValue)
	})
}

func TestFetchSheets(t *testing.T) {
	_, err := FetchSheets(srv.Service, spreadsheetId)

	assert.Nil(t, err)
}
