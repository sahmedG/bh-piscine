package piscine

func Split(s, sep string) []string {
	length := 0
	sublength := 0
	for i := range s {
		length = i + 1
	}
	for i := range sep {
		sublength = i + 1
	}
	for i := 0; i < sublength; i++ {
		s += " "
	}
	prev := false
	len := 0
	for i := 0; i < length; i++ {
		if (s[i:i+sublength] == sep) && !prev {
			prev = true
			len++
		} else {
			prev = false
		}
	}
	len++
	arr := make([]string, len)
	word := ""
	arindex := 0
	for i := 0; i < length; i++ {
		if s[i:i+sublength] == sep {
			l := 0
			for i := range word {
				l = i + 1
			}
			if l == 0 {
				continue
			}
			arr[arindex] = word
			arindex++
			word = ""
			i = i + sublength - 1
			continue
		}
		word += string(s[i])
	}
	l := 0
	for i := range word {
		l = i + 1
	}
	if l != 0 {
		arr[arindex] = word
	}
	return arr
}
