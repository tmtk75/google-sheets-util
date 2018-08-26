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
		// Print columns A and E, which correspond to indices 0 and 4.
		//fmt.Printf("%s\n", row)
		e := make(map[string]interface{})
		for j, v := range row {
			idx := k[j]
			if a, b := v.(string); b {
				var t map[string]interface{}
				//fmt.Printf("a: %v\n", a)
				if err := json.Unmarshal([]byte(a), &t); err == nil {
					e[idx] = t // as JSON
				} else {
					var s interface{}
					if err := json.Unmarshal([]byte(a), &s); err == nil {
						e[idx] = s // as something
					} else {
						if p, err := strconv.ParseBool(a); err == nil {
							e[idx] = p
						} else {
							if a == "<null>" {
								fmt.Printf("hi: %v, %v, %v\n", i, j, a)
								e[idx] = nil
							} else {
								e[idx] = a // as raw string
							}
						}
					}
				}
			} else {
				fmt.Printf("fail to assert\n")
				e[k[j]] = v
			}
		}
		m[i-1] = e
	}
	return m, nil
}
