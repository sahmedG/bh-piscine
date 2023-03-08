package main

import "fmt"


func Split(s, sep string) []string {
	word := ""
	var str_arr []string
	if len(s) < 0 {
		return str_arr
	}
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(sep)-1; j++ {
			if s[i] == sep[j] && s[i+1] == sep[j+1] {
				i++
				if word == "" {
					continue
				}
				str_arr = append(str_arr, word)
				word = ""
			} else {
				word += string(s[i])
			}
		}
	}
	word += string(s[len(s)-1])
	str_arr = append(str_arr, word)
	return str_arr
}

func main() {
	s := "HelloHAWhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
