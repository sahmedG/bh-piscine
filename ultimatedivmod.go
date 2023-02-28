package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	var x = *a
	var y = *b
	*a = *a / *b
	var t = x % y
	*b = *&t
}
