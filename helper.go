package main

import (
	"strings"
)

var cityCountryList = map[string]string{
	"berlin": "de",
	"cairo":  "eg",
}

func GetCountryByCityName(s string) (string, bool) {
	name := strings.ToLower(s)
	value, exist := cityCountryList[name]
	return value, exist
}
