package utils

import (
	"encoding/json"
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

// JSONBytesEqual compares the JSON in two byte slices.
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	jsonConfig := jsoniter.ConfigCompatibleWithStandardLibrary
	if err := jsonConfig.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := jsonConfig.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

// ToRawMessage is a converter any interface to RawMessage
func ToRawMessage(val interface{}) *json.RawMessage {
	jsonConfig := jsoniter.ConfigCompatibleWithStandardLibrary
	valAsByte, _ := jsonConfig.Marshal(val)

	return (*json.RawMessage)(&valAsByte)
}
