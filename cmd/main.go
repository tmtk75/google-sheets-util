package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tmtk75/google-sheets-util"
)

func main() {
	srv, err := sheetsutil.NewSheetService()
	if err != nil {
		log.Fatalf("%v", err)
	}

	// https://docs.google.com/spreadsheets/d/13zo1qomUg6Dgh0B8Sr53BwhFnqrsxcYnso_p9pM9meg/edit#gid=0
	spreadsheetId := "13zo1qomUg6Dgh0B8Sr53BwhFnqrsxcYnso_p9pM9meg"
	res, err := srv.Spreadsheets.Values.Get(spreadsheetId, "nested!A1:A").Do()
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
