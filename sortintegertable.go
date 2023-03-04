package piscine

func SortIntegerTable(table []int) {
	var temp int
	for i := 0; i <= len(table)-1; i++ {
		for j := i + 1; j <= len(table)-1; j++ {
			if table[i] > table[j] {
				temp = table[i]
				table[i] = table[j]
				table[j] = temp
			}
		}
	}
}
