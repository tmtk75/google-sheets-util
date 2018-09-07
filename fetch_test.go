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
	sheets, err := FetchSheets(srv.Service, spreadsheetId)

	assert.Nil(t, err)
	assert.Equal(t, 4, len(sheets))
	assert.Equal(t, "capital", sheets[0].Title)
	assert.Equal(t, "nested", sheets[1].Title)
	assert.Equal(t, "typed", sheets[2].Title)
	assert.Equal(t, "sparse", sheets[3].Title)

	assert.Equal(t, int64(0), sheets[0].SheetId)
	assert.Equal(t, int64(1054988972), sheets[1].SheetId)
	assert.Equal(t, int64(2135743257), sheets[2].SheetId)
	assert.Equal(t, int64(291192554), sheets[3].SheetId)

}
