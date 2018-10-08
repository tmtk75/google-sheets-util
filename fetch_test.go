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
	assert.Equal(t, 5, len(sheets))
	assert.Equal(t, "capital", sheets[0].Title)
	assert.Equal(t, "nested", sheets[1].Title)
	assert.Equal(t, "typed", sheets[2].Title)
	assert.Equal(t, "sparse", sheets[3].Title)
	assert.Equal(t, "hiddenByUser", sheets[4].Title)

	assert.Equal(t, int64(0), sheets[0].SheetId)
	assert.Equal(t, int64(1054988972), sheets[1].SheetId)
	assert.Equal(t, int64(2135743257), sheets[2].SheetId)
	assert.Equal(t, int64(291192554), sheets[3].SheetId)
	assert.Equal(t, int64(167016578), sheets[4].SheetId)
}

func TestFetchSheet(t *testing.T) {
	t.Run("hiddenByUser", func(t *testing.T) {
		sheet, err := FetchSheet(srv.Service, spreadsheetId, "hiddenByUser!A1:C")
		assert.Nil(t, err)
		// GridData row
		assert.Equal(t, false, sheet.GridData.RowMetadata[0].HiddenByUser)
		assert.Equal(t, true, sheet.GridData.RowMetadata[1].HiddenByUser)
		assert.Equal(t, false, sheet.GridData.RowMetadata[2].HiddenByUser)
		assert.Equal(t, true, sheet.GridData.RowMetadata[3].HiddenByUser)
		assert.Equal(t, false, sheet.GridData.RowMetadata[4].HiddenByUser)
		// GridData column
		assert.Equal(t, false, sheet.GridData.ColumnMetadata[0].HiddenByUser)
		assert.Equal(t, true, sheet.GridData.ColumnMetadata[1].HiddenByUser)
		assert.Equal(t, false, sheet.GridData.ColumnMetadata[2].HiddenByUser)
		// Cells
		assert.Equal(t, "1", sheet.Cells[0][0].FormattedValue)
		assert.Equal(t, "2", sheet.Cells[1][0].FormattedValue)
		assert.Equal(t, "3", sheet.Cells[2][0].FormattedValue)
		assert.Equal(t, "4", sheet.Cells[3][0].FormattedValue)
		assert.Equal(t, "5", sheet.Cells[4][0].FormattedValue)
		assert.Equal(t, "A", sheet.Cells[0][1].FormattedValue)
		assert.Equal(t, "B", sheet.Cells[1][1].FormattedValue)
		assert.Equal(t, "C", sheet.Cells[2][1].FormattedValue)
		assert.Equal(t, "D", sheet.Cells[3][1].FormattedValue)
		assert.Equal(t, "E", sheet.Cells[4][1].FormattedValue)
		assert.Equal(t, "a", sheet.Cells[0][2].FormattedValue)
		assert.Equal(t, "b", sheet.Cells[1][2].FormattedValue)
		assert.Equal(t, "c", sheet.Cells[2][2].FormattedValue)
		assert.Equal(t, "d", sheet.Cells[3][2].FormattedValue)
		assert.Equal(t, "e", sheet.Cells[4][2].FormattedValue)
	})
}
