package sheetsutil

import sheets "google.golang.org/api/sheets/v4"

func FetchCells(srv *sheets.Service, id, address string) ([][]*sheets.CellData, error) {
	res, err := srv.Spreadsheets.Get(id).Ranges(address).Fields("sheets(properties,data.rowData.values(formattedValue,note))").Do()
	if err != nil {
		return nil, err
	}
	return ToCellData2D(res), nil
}

func ToCellData2D(res *sheets.Spreadsheet) [][]*sheets.CellData {
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
