package piscine

func UltimateDivMod(a *int, b *int) {
	var x = *a
	var y = *b
	*a = x / y
	*b = x % y
}
