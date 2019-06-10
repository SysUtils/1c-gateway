package odata

import "encoding/json"

// Type converter for any types through conversion to JSON and back
func ConvertType(in interface{}, out interface{}) error {
	data, err := json.Marshal(in)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, out)
	return nil
}
