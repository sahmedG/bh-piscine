package piscine

func MakeRange(min, max int) []int {
	len := max - min
	if len > 0 {
		slice := make([]int, len)
		for j := min; j < max; j++ {
			slice[j-min] = j
		}
		return slice
	}
	slice := make([]int, 0)
	return slice
}
