package sheetsutil

import (
	"strconv"

	"google.golang.org/api/sheets/v4"
)

func ToMap(resp *sheets.ValueRange) ([]map[string]interface{}, error) {
	k := make([]string, len(resp.Values[0]))
	m := make([]map[string]interface{}, len(resp.Values)-1)
	for i, row := range resp.Values {
		if i == 0 {
			for j, v := range row {
				if e, b := v.(string); b {
					k[j] = e
				}
			}
			//fmt.Printf("header: %v\n", k)
			continue
		}
		// Print columns A and E, which correspond to indices 0 and 4.
		//fmt.Printf("%s\n", row)
		e := make(map[string]interface{})
		for j, v := range row {
			if s, b := v.(string); b {
				if n, err := strconv.ParseFloat(s, 32 /*bit*/); err == nil {
					e[k[j]] = n
				} else {
					e[k[j]] = v
				}
			}
		}
		m[i-1] = e
	}
	return m, nil
}
