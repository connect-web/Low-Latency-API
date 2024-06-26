package util

import (
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
