package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func convertInttoRune(nbr int) {
	count := '0'
	for i := 1; i <= nbr%10; i++ {
		count++
	}
	if nbr/10 != 0 {
		convertInttoRune(nbr / 10)
	}
	z01.PrintRune(count)
}

func main() {
	points := &point{}
	setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	convertInttoRune(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	convertInttoRune(points.y)
	z01.PrintRune('\n')
}
