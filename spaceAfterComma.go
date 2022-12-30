package main

import (
	"fmt"
	"regexp"
)

func insertCommaSpaces(s string) string {
	replace := regexp.MustCompile(",")
	return replace.ReplaceAllString(s, ", ")
}

func main() {
	input := "Q1 months are,January,February,March."
	output := insertCommaSpaces(input)
	fmt.Println(output)
}
