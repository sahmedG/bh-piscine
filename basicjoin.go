package piscine

func BasicJoin(elems []string) string {
	words := ""
	i := 0
	for range elems {
		words = words + elems[i]
		i++
	}
	return words
}
