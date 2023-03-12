package piscine

func Any(f func(string) bool, a []string) bool {
	var result bool
	for i := range a {
		result = f(a[i])
		if result == true {
			return result
		}
	}
	return result
}
