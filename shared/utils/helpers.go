package utils

import (
	"encoding/json"
	"strings"
)

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func StringToPointer(value string) *string {
	return &value
}

func BoolToPointer(value bool) *bool {
	return &value
}

func IntToPointer(value int) *int {
	return &value
}

func Float64ToPointer(value float64) *float64 {
	return &value
}

func HasRecord(err error) bool {
	index := strings.Index(err.Error(), "violates unique constraint")

	return index > 0
}

func ForeignKeyNotExist(err error) bool {
	index := strings.Index(err.Error(), "violates foreign key constraint")

	return index > 0
}

func StructToInterface(value interface{}) map[string]interface{} {
	in := make(map[string]interface{})

	updateData, _ := json.Marshal(value)
	json.Unmarshal(updateData, &in)

	return in
}

func HasValue(phrase, value string) bool {
	index := strings.Index(phrase, value)

	return index > 0
}
