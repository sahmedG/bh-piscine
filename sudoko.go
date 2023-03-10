package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Args := os.Args[1:]
	arguments := []string(Args)

	// check the validity of the inputs and create solution table
	if checkInput(arguments) == true {
		table := [9][9]rune{}
		table = fillTable(table, arguments)

	// upon recieving a slution fill-out the created table
		if isSolve(&table) == true {
			for y := 0; y < 9; y++ {
				for x := 0; x < 9; x++ {
					if x != 8 {
						z01.PrintRune(rune(table[y][x]))
						z01.PrintRune(32)
					} else {
						z01.PrintRune(rune(table[y][x]))
					}
				}
				z01.PrintRune(10)
			}
		} else {
			fmt.Println("Error") // if the input in not valid raise an error
		}
	}
}

// input validity check
func checkInput(args []string) bool {
	// check if the input array has 9 elements
	if len(args) != 9 {
		fmt.Println("Error") // Invalid input
		return false
	}
	// check each element inside the array has 9 elemnts
	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			fmt.Println("Error") //  invalid input
			return false
		}
	}
	// check invalid values 
	for i := 0; i < len(args); i++ {
		for _, value := range args[i] {
			if value < 49 && value > 57 {
				fmt.Println("Error") // IInvalid input
				return false
			}
		}
	}
	return true
}

// initiate the created table with provided input
func fillTable(table [9][9]rune, args []string) [9][9]rune {
	for i := range args {
		for j := range args[i] {
			table[i][j] = rune(args[i][j])
		}
	}
	return table
}

// checking for empty spot in the table
func isDots(table *[9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == '.' {
				return true
			}
		}
	}
	return false
}

// check if it fits

func isValid(table *[9][9]rune, x int, y int, z rune) bool {
	// check double int
	for i := 0; i < 9; i++ {
		if z == table[i][x] {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if z == table[y][j] {
			return false
		}
	}
	//square check
	a := x / 3
	b := y / 3
	for k := 3 * a; k < 3*(a+1); k++ {
		for l := 3 * b; l < 3*(b+1); l++ {
			if z == table[l][k] {
				return false
			}
		}
	}
	return true
}
//backtracking solver
func isSolve(table *[9][9]rune) bool {

	//check no more empty spots in the table
	if !isDots(table) {
		return true
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if table[y][x] == '.' {
				for z := '1'; z <= '9'; z++ {
					if isValid(table, x, y, z) {
						table[y][x] = z
						if isSolve(table) {
							return true
						}
					}
					table[y][x] = '.'
				}
				return false
			}
		}
	}
	return false
}
