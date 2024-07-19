package util

import (
	"encoding/json"
	"strconv"
	"strings"
)

// ConvertStringArrayToInterfaceArray converts a slice of strings to a slice of interfaces
func ConvertStringArrayToInterfaceArray(stringArray []string) []interface{} {
	var interfaceArray []interface{} = make([]interface{}, len(stringArray))
	for i, v := range stringArray {
		interfaceArray[i] = v
	}
	return interfaceArray
}

func StringToIntText(text string) (int, error) {
	text = strings.ToLower(text)
	if strings.Contains(text, "m") {
		text = strings.ReplaceAll(text, "m", "000000")
	}
	if strings.Contains(text, "k") {
		text = strings.ReplaceAll(text, "k", "000")
	}
	number, err := strconv.Atoi(text)
	return number, err
}

func DecodeJSONToInt64Map(data []byte) map[string]int64 {
	var result map[string]int64
	if err := json.Unmarshal(data, &result); err != nil {
		return make(map[string]int64)
	}
	return result
}

func DecodeJSONToIntMap(data []byte) map[string]int {
	var result map[string]int
	if err := json.Unmarshal(data, &result); err != nil {
		return make(map[string]int)
	}
	return result
}

func DecodeJSONToFloat64Map(data []byte) map[string]float64 {
	var result map[string]float64
	if err := json.Unmarshal(data, &result); err != nil {
		return make(map[string]float64)
	}
	return result
}
