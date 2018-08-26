package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/tmtk75/google-sheets-util"
)

func main() {
	var (
		sheetId    = flag.String("spreadsheet-id", "13zo1qomUg6Dgh0B8Sr53BwhFnqrsxcYnso_p9pM9meg", "Spreadsheet ID")
		sheetName  = flag.String("sheet-name", "nested", "Sheet name")
		sheetRange = flag.String("sheet-range", "A1:A", "Range")
		path       = flag.String("credential-path", "./credentials.json", "Path to credentials in JWT")
	)
	flag.Parse()

	srv, err := sheetsutil.NewSheetService(*path)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// https://docs.google.com/spreadsheets/d/13zo1qomUg6Dgh0B8Sr53BwhFnqrsxcYnso_p9pM9meg/edit#gid=0
	spreadsheetId := *sheetId
	address := fmt.Sprintf("%v!%v", *sheetName, *sheetRange)
	res, err := srv.Spreadsheets.Values.Get(spreadsheetId, address).Do()
	if err != nil {
		log.Fatalf("Spreadsheets.Values.Get: %v", err)
	}

	rows, err := sheetsutil.ToMap(res)
	if err != nil {
		log.Fatalf("ToMap: %v", err)
	}

	b, err := json.Marshal(rows)
	if err != nil {
		log.Fatalf("ToMap: %v", err)
	}

	fmt.Println(string(b))
}
