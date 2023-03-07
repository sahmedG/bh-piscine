package piscine

func ConcatParams(args []string) string {
	string1 := ""
	var str_arr []string
	if len(args) > 0 {
		for i := 0; i <= len(args)-1; i++ {
			str_arr = append(str_arr, string(args[i]))
			str_arr = append(str_arr, "\n")
		}
	}
	for j := 0; j < len(str_arr)-1; j++ {
		string1 = string1 + str_arr[j]
	}
	return string1
}
