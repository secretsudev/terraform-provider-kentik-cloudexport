package provider

import "fmt"

func getObjectFromNestedResourceData(data interface{}) (map[string]interface{}, error) {
	dataSlice, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data type, got: %T, want: []interface{}", data)
	}

	if len(dataSlice) == 0 {
		return nil, nil
	}

	if dataSlice[0] == nil {
		return nil, nil
	}

	m, ok := dataSlice[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf(
			"invalid dataSlice[0] type, got: %T, want: map[string]interface{}",
			dataSlice[0],
		)
	}

	return m, nil
}
