package piscine

func Index(s string, toFind string) int {
	l := len(s) - len(toFind)
	count := 0
	i := 0
	for i = 0; i <= l; i++ {
		for j := 0; j < len(toFind); j++ {
			if s[i] == toFind[j] {
				count++
			}
			if count == len(toFind) {
				break
			}
		}
		if count == len(toFind) {
			break
		}
	}
	//fmt.Println(count)
	if count == len(toFind) {
		for i = 0; i < len(toFind); i++ {
			for j := 0; j <= l; j++ {
				if s[j] == toFind[i] {
					return j
				}
			}
		}
	}
	return -1
}
