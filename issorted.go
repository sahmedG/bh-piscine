package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	res := 0
	count := 0
	for i := 0; i < len(a)-1; i++ {
		res = f(a[i], a[i+1])
		if res == 1 {
			count += 1
		}
	}
	if count == len(a)-1 {
		return true
	}
	return false
}

func SortIntegerTable(a, b int) int {
	if a > b {
		return -1
	} else if a == b {
		return 0
	} else {
		return 1
	}
}
