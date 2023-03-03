package main

import "fmt"

func main() {

	height := 3
	width := 5

	for h := 1; h <= height; h++ {
		for w := 1; w <= width; w++ {
			if h == 1 && w == 1 {
				fmt.Print("o")
			}
			if w == 1 || w == width {
				if h > 1 && h < height {
					fmt.Print("|")
				} else if w == width {
					fmt.Print("o")
				} else if w == 1 && h == height {
					fmt.Print("o")
				} else {
					fmt.Print("-")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}