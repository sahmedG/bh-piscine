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
	xprint := []rune{'x', ' ', '=', ' '}
	points := &point{}
	setPoint(points)
	for _, value := range xprint {
		z01.PrintRune(value)
	}
	convertInttoRune(points.x)
	yprint := []rune{',', ' ', 'y', ' ', '=', ' '}
	for _, value := range yprint {
		z01.PrintRune(value)
	}
	convertInttoRune(points.y)
	z01.PrintRune('\n')
}
