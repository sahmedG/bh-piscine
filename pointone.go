package piscine

import "fmt"

func PointOne(n *int) {
	*n = *n + 1
}

func main() {
	n := 0
	PointOne(&n)
	fmt.Println(n)
}
