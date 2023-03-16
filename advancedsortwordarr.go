package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if f(a[i], a[j]) < 0 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
