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
