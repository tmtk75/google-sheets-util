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
				nv, err := ToValue(a, i, j)
				if err != nil {
					return nil, err
				}
				e[idx] = nv
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

type ValueFunc func(a string, i, j int) (interface{}, error)

/*
 * Table to MapArray.
 * 1st line is used as key names.
 *
 * [0] a | b | c
 *     --|---|---  --> [{a:1, b:2, c:3}, {a:A, b:B, c:C}]
 * [1] 1 | 2 | 3
 * [2] A | B | C
 */
func ToMapArray(vals [][]string, toValue ValueFunc) ([]map[string]interface{}, error) {
	k := make([]string, len(vals[0]))
	m := make([]map[string]interface{}, len(vals)-1)
	for i, row := range vals {
		if i == 0 {
			for j, v := range row {
				k[j] = v
			}
			continue
		}

		e := make(map[string]interface{})
		for j, v := range row {
			idx := k[j]
			tv, err := toValue(v, i, j)
			if err != nil {
				return nil, err
			}
			e[idx] = tv
		}

		m[i-1] = e
	}
	return m, nil
}

func ToValue(a string, i, j int) (interface{}, error) {
	var t map[string]interface{}
	if err := json.Unmarshal([]byte(a), &t); err == nil {
		return t, nil // as JSON
	}

	var s interface{}
	if err := json.Unmarshal([]byte(a), &s); err == nil {
		return s, nil // as something
	}

	if p, err := strconv.ParseBool(a); err == nil {
		return p, nil
	}

	return a, nil // as raw string
}
