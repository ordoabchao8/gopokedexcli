package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var slicey []string
	fields := strings.Fields(text)
	for _, f := range fields {
		slicey = append(slicey, strings.ToLower(f))
	}
	return slicey
}