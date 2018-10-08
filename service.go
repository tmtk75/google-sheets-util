package sheetsutil

import (
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

type Service struct {
	*sheets.Service
}

func NewSheetService(path string) (*Service, error) {
	b, err := ioutil.ReadFile(path) // TODO: Support anonymous access, can I?
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		return nil, err
	}

	client := conf.Client(oauth2.NoContext)

	srv, err := sheets.New(client)
	if err != nil {
		return nil, err
	}

	return &Service{Service: srv}, nil
}

func (s *Service) FetchSheet(id, address string) (*Sheet, error) {
	return FetchSheet(s.Service, id, address)
}
