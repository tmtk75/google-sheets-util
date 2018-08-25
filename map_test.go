package sheetsutil

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

const spreadsheetId = "13zo1qomUg6Dgh0B8Sr53BwhFnqrsxcYnso_p9pM9meg"

var (
	client *http.Client
	srv    *sheets.Service
)

func init() {
	b, err := ioutil.ReadFile("./credentials.json") // TODO: Support anonymous access, can I?
	if err != nil {
		panic(err)
	}
	conf, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		panic(err)
	}

	client = conf.Client(oauth2.NoContext)

	srv, err = sheets.New(client)
	if err != nil {
		panic(err)
	}
}

func TestToMap(t *testing.T) {
	// https://docs.google.com/spreadsheets/d/13zo1qomUg6Dgh0B8Sr53BwhFnqrsxcYnso_p9pM9meg/edit#gid=0

	t.Run("properties", func(t *testing.T) {
		sheet, err := srv.Spreadsheets.Get(spreadsheetId).Do()
		if err != nil {
			t.Fatalf("Spreadsheets.Get: %v", err)
		}

		expect := "github.com/tmtk75/google-sheets-util"
		if sheet.Properties.Title != expect {
			t.Fatalf("Title: expected: %v, actual: %v", expect, sheet.Properties.Title)
		}
	})

	t.Run("capital", func(t *testing.T) {
		res, err := srv.Spreadsheets.Values.Get(spreadsheetId, "capital!A1:B").Do()
		if err != nil {
			t.Fatalf("Spreadsheets.Values.Get: %v", err)
		}

		rows, err := ToMap(res)
		if err != nil {
			t.Fatalf("Spreadsheets.Values.Get: %v", err)
		}

		country := []string{"Japan", "China", "Korea"}
		capital := []string{"Tokyo", "Beijin", "Soul"}
		for i, r := range rows {
			if r["country"] != country[i] {
				t.Fatalf("country[%d]: expected: %v, actual: %v", i, country[i], r["country"])
			}
			if r["capital"] != capital[i] {
				t.Fatalf("capital[%d]: expected: %v, actual: %v", i, capital[i], r["capital"])
			}
		}
	})
}

func TestToMapNested(t *testing.T) {
	res, err := srv.Spreadsheets.Values.Get(spreadsheetId, "nested!A1:B").Do()
	if err != nil {
		t.Fatalf("Spreadsheets.Values.Get: %v", err)
	}

	rows, err := ToMap(res)

	if e, b := rows[0]["body"].(map[string]interface{}); !b {
		it := rows[0]["body"]
		t.Fatalf(`failed type assert as rows[0]["body"]: type: %v, %v`, reflect.TypeOf(it), it)
	} else {
		if e["a"] != 1 {
			t.Fatalf("a: expected: %v, actual: %v", 1, e["a"])
		}
		if e["b"] != "xxx" {
			t.Fatalf("b: expected: %v, actual: %v", "xxx", e["a"])
		}
		if e["c"] != true {
			t.Fatalf("c: expected: %v, actual: %v", true, e["a"])
		}
		if e["d"] != nil {
			t.Fatalf("d: expected: %v, actual: %v", nil, e["a"])
		}
	}
}
