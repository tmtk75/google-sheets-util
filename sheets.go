package sheetsutil

import (
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func NewSheetService(path string) (*sheets.Service, error) {
	b, err := ioutil.ReadFile(path) // TODO: Support anonymous access, can I?
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}

	client := conf.Client(oauth2.NoContext)

	srv, err := sheets.New(client)
	if err != nil {
		return nil, err
	}

	return srv, nil
}

func Fetch(srv *sheets.Service, id, address string) ([][]*sheets.CellData, error) {
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
