package piscine

func Capitalize(s string) string {
	str_rune := []rune(s)
	i := 0
	is_letter := false
	is_digit := false

	for range str_rune {
		if (str_rune[i] >= 65 && str_rune[i] <= 90) || (str_rune[i] >= 97 && str_rune[i] <= 122) || (str_rune[i] >= 48 && str_rune[i] <= 57) {
			if is_letter == false {
				is_digit = true
				is_letter = true
			} else {
				is_digit = false
			}
		} else {
			is_digit = false
			is_letter = false
		}
		if is_letter == true {
			if is_digit == false {
				if str_rune[i] >= 65 && str_rune[i] <= 90 {
					str_rune[i] += 32
				}
			} else if str_rune[i] >= 97 && str_rune[i] <= 122 {
				str_rune[i] -= 32
			}
		}
		i++
	}
	return string(str_rune)
}
