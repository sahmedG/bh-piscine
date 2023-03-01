package piscine

func UltimateDivMod(a *int, b *int) {
	var x = *a
	var y = *b
	var k = x / y
	var t = x % y
	*b = *&t
	*a = *&k
}
