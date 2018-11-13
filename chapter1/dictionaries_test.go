package chapter1

import (
	"encoding/json"
	"errors"
	"testing"
)

// TestDcitionaries1 is a testcase about how marshal/unmarshal maps
func TestDcitionaries1(t *testing.T) {
	// Unmarshal a JSON string to a map
	t.Logf("Unmarshal a JSON string to a map")
	jsonStr := `{"field1": 10, "field2": "hello world", "field3": true, "field4": "2018-11-13T02:23:28+02:00Z"}`
	m1 := make(map[string]interface{})
	t.Logf("JSON String: %+v", jsonStr)
	t.Logf("map before update: %+v", m1)
	err := json.Unmarshal([]byte(jsonStr), &m1)
	if err != nil {
		t.Fail()
	}
	t.Logf("map after update: %+v", m1)
	t.Logf("\n")

	// Marshal a map to JSON string
	t.Logf("Marshal a map to JSON string")
	jsonStr = ``
	m1["field1"] = m1["field1"].(float64) + 10
	m1["field2"] = m1["field2"].(string) + "!!!"
	m1["field3"] = !m1["field3"].(bool)
	t.Logf("Modified map: %+v", m1)
	jsonBytes, err := json.Marshal(m1)
	if err == nil {
		err = errors.New("just testing")
	}
	if err != nil {
		t.Logf("Caught fake error '%+v'", err)
	}
	t.Logf("JSON String: %+v", string(jsonBytes))
	t.Logf("\n")
}
