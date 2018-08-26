package sheetsutil

import (
	"reflect"
	"testing"

	"google.golang.org/api/sheets/v4"
)

const spreadsheetId = "13zo1qomUg6Dgh0B8Sr53BwhFnqrsxcYnso_p9pM9meg"

var (
	srv *sheets.Service
)

func init() {
	s, err := NewSheetService("./credentials.json")
	if err != nil {
		panic(nil)
	}
	srv = s
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
	if err != nil {
		t.Fatalf("ToMap: %v", err)
	}

	if e, b := rows[0]["body"].(map[string]interface{}); !b {
		it := rows[0]["body"]
		t.Fatalf(`failed type assert for rows[0]["body"]: type: %v, %v`, reflect.TypeOf(it), it)
	} else {
		if e["a"] != float64(1) {
			t.Fatalf("a: expected: %v, actual: %v", float64(1), e["a"])
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

func TestToMapTyped(t *testing.T) {
	res, err := srv.Spreadsheets.Values.Get(spreadsheetId, "typed!A1:A").Do()
	if err != nil {
		t.Fatalf("Spreadsheets.Values.Get: %v", err)
	}

	rows, err := ToMap(res)
	if err != nil {
		t.Fatalf("ToMap: %v", err)
	}

	expect := []interface{}{float64(123), true, false, nil, "hello"}
	for i, e := range rows {
		if e["v"] != expect[i] {
			t.Fatalf("[%v]: expected: %v(%v), actual: %v(%v)",
				i, expect[i], reflect.TypeOf(expect[i]),
				e["v"], reflect.TypeOf(e["v"]))
		}
	}
}

func TestToMapSparse(t *testing.T) {
	res, err := srv.Spreadsheets.Values.Get(spreadsheetId, "sparse!A1:D").Do()
	if err != nil {
		t.Fatalf("Spreadsheets.Values.Get: %v", err)
	}

	rows, err := ToMap(res)
	if err != nil {
		t.Fatalf("ToMap: %v", err)
	}

	expect := []map[string]interface{}{
		{"a": float64(1), "b": nil, "c": nil, "d": nil},
		{"a": "", "b": float64(2), "c": nil, "d": nil},
		{"a": "", "b": "", "c": float64(3), "d": nil},
		{"a": "", "b": "", "c": "", "d": float64(4)},
	}
	for i, e := range rows {
		for j, v := range expect[i] {
			idx := j
			if e[idx] != v {
				t.Fatalf("[%v, %v]: expected: %v, actual: %v", j, i, v, e[idx])
			}
		}
	}
}
