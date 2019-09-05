package utils

import (
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

func HasRecord(err error) bool {
	index := strings.Index(err.Error(), "violates unique constraint")

	return index > 0
}

func ForeignKeyNotExist(err error) bool {
	index := strings.Index(err.Error(), "violates foreign key constraint")

	return index > 0
}
