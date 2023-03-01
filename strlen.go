package piscine

func StrLen(s string) int {
	l := utf8.RuneCountInString(s)
	return l
}
