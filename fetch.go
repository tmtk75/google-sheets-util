package sheetsutil

import (
	"fmt"

	sheets "google.golang.org/api/sheets/v4"
)

// simple wrapper struct
type Sheet struct {
	*sheets.SheetProperties
	Cells    [][]*sheets.CellData // may be nil
	GridData *sheets.GridData     // may be nil
}

func ToCellDataArrays(res *sheets.Spreadsheet) [][]*sheets.CellData {
	rows := make([][]*sheets.CellData, 0)
	for _, sheet := range res.Sheets {
		for _, d := range sheet.Data {
			for _, e := range d.RowData {
				l := make([]*sheets.CellData, len(e.Values))
				for i, v := range e.Values {
					l[i] = v
				}
				rows = append(rows, l)
			}
		}
	}
	return rows
}

func FetchSheets(srv *sheets.Service, id string) ([]Sheet, error) {
	res, err := srv.Spreadsheets.Get(id).Fields("sheets(properties)").Do()
	if err != nil {
		return nil, err
	}
	//fmt.Printf("%v", res.Sheets[0].Properties)
	sheets := make([]Sheet, len(res.Sheets))
	for i, s := range res.Sheets {
		sheets[i] = Sheet{SheetProperties: s.Properties}
	}
	return sheets, nil
}

func FetchSheet(srv *sheets.Service, id, address string) (*Sheet, error) {
	res, err := srv.Spreadsheets.Get(id).Ranges(address).Fields("sheets(properties,data(rowMetadata.hiddenByUser,columnMetadata.hiddenByUser,rowData.values(formattedValue,note)))").Do()
	if err != nil {
		return nil, err
	}

	if len(res.Sheets) != 1 {
		return nil, fmt.Errorf("unexpected count of sheets. len: %v", len(res.Sheets))
	}
	it := res.Sheets[0]

	if len(it.Data) != 1 {
		return nil, fmt.Errorf("unexpected count of data. len: %v", len(it.Data))
	}

	// Fill Cells
	s := Sheet{
		SheetProperties: it.Properties,
		GridData:        it.Data[0],
		Cells:           ToCellDataArrays(res),
	}

	return &s, nil
}
