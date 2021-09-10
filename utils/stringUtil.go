package utils

import (
	"regexp"
	"strconv"
)

func GetStreamData(str string) float64 {
	carry := 1.0
	k, _ := regexp.Compile("[0-9]*K")
	if k.Match([]byte(str)) {
		str = str[0 : len(str)-1]
		carry *= 1000
	}
	m, _ := regexp.Compile("[0-9]*M")
	if m.Match([]byte(str)) {
		str = str[0 : len(str)-1]
		carry *= 1000000
	}
	g, _ := regexp.Compile("[0-9]*G")
	if g.Match([]byte(str)) {
		str = str[0 : len(str)-1]
		carry *= 1000000
	}
	num, _ := strconv.ParseFloat(str, 64)
	return num * carry
}
