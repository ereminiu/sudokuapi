package models

import (
	"log"
)

const (
	CorrectGrid      = 1
	UnvalidRow       = 2
	UnvalidCol       = 3
	UnvalidRowAmount = 4
	UnvalidColAmount = 5
)

func Decode(code int) string {
	if code == 1 {
		return "Grid is correct"
	} else if code == 2 {
		return "Error: two same values in the same row"
	} else if code == 3 {
		return "Error: two same values in the same column"
	} else if code == 4 {
		return "Error: Unvalid amount of rows, it should be 9"
	} else if code == 5 {
		return "Error: Unvalid amount of columns, it should be 9"
	}

	log.Fatalf("Unexpected code %d\n", code)
	return ""
}
