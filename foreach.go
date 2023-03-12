package piscine

import "fmt"

func ForEach(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		PrintNbr(a[i])
	}
	fmt.Println()
}
