package common

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func StrToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Cannot convert to int", str)
	}

	return value
}

func SubtractStr(str1 string, str2 string) string {
	var unique []string

	for _, char := range str1 {
		if !strings.Contains(str2, string(char)) {
			unique = append(unique, string(char))
		}
	}

	return strings.Join(unique, "")
}

func IntersectStr(str1 string, str2 string) string {
	var intersection []string

	for _, char := range str1 {
		if strings.Contains(str2, string(char)) {
			intersection = append(intersection, string(char))
		}
	}

	return strings.Join(intersection, "")
}

func ContainsInAnyOrder(str string, chars string) bool {
	for _, char := range chars {
		if !strings.Contains(str, string(char)) {
			return false
		}
	}

	return true
}

func IsLowerCased(str string) bool {
	for _, char := range str {
		if !unicode.IsLower(char) {
			return false
		}
	}

	return true
}

func Update(str string, ix int, value rune) string {
	return str[0:ix] + string(value) + str[ix+1:]
}
