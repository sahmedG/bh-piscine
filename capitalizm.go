package piscine

func capitalizeByteSlice(str string) string {
	//bs := []string(str)
	if len(str) == 0 {
		return ""
	}

	string1 := ""
	final_string := ""
	var str_arr []string

	for i := 0; i <= len(str)-1; i++ {

		if str[i] >= 'A' && str[i] <= 'Z' || str[i] >= 'a' && str[i] <= 'z' || str[i] >= '0' && str[i] <= '9' {
			string1 = string1 + string(str[i])
		}
		if !(str[i] >= 'A' && str[i] <= 'Z' || str[i] >= 'a' && str[i] <= 'z' || str[i] >= '0' && str[i] <= '9') {
			string1 = string1 + string(str[i])
			str_arr = append(str_arr, string1)
			string1 = ""
		}
	}
	fmt.Println(str_arr)

	for t := 0; t < len(str_arr); t++ {
		for k := 0; k < len(str_arr[t]); k++ {
			if str_arr[t][0] >= 97 && str_arr[t][0] <= 122 {
				chr := str_arr[t][0] - 32
				final_string = string(chr) + string(str_arr[t][k])
			} else if str_arr[t][k] >= 64 && str_arr[t][k] <= 90 {
				chr := str_arr[t][k] + 32
				final_string = string(str_arr[t][k]) + string(chr)
			}
		}
	}
	return final_string
}
