package piscine

import "unicode/utf8"

func StrLen(s string) int {
	l := utf8.RuneCountInString(s)
	return l
}
