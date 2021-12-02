package common

import (
	"log"
	"strconv"
)

func StrToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Cannot convert to int", str)
	}

	return value
}
