package sheetsutil

import (
	"encoding/json"
	"fmt"
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

		e := make(map[string]interface{})
		for j, v := range row {
			idx := k[j]
			if a, b := v.(string); b {
				var t map[string]interface{}
				if err := json.Unmarshal([]byte(a), &t); err == nil {
					e[idx] = t // as JSON
					continue
				}

				var s interface{}
				if err := json.Unmarshal([]byte(a), &s); err == nil {
					e[idx] = s // as something
					continue
				}

				if p, err := strconv.ParseBool(a); err == nil {
					e[idx] = p
					continue
				}

				e[idx] = a // as raw string
				continue
			} else {
				return nil, fmt.Errorf("unexpected path: %v", row)
			}

			e[idx] = v
		}

		m[i-1] = e
	}
	return m, nil
}
