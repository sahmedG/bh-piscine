package piscine

func MakeRange(min, max int) []int {
	if max > min {
		len := max - min
		slice := make([]int, len)
		for j := min; j < max; j++ {
			slice[j-min] = j
		}
		return slice
	}
	var slice []int
	return slice
}
