package main

import (
	"fmt"
	"os"
)

func main() {
	Args := os.Args[1:]
	arguments := []string(Args)

	// check the validity of the inputs and create solution table

	if ValidateInput(arguments) == true {
		table := [9][9]rune{}
		// create a board with provided elements
		table = fillTable(table, arguments)

		// upon recieving a solution fill-out the created table
		if IsItSolved(&table) == true {
			if isBoardValid(&table) {
				for y := 0; y < 9; y++ {
					for x := 0; x < 9; x++ {
						if x != 8 {
							fmt.Print(string(table[y][x]))
							fmt.Print(" ")
						} else {
							fmt.Print(string(table[y][x]))
						}
					}
					fmt.Println()
				}

			} else {
				fmt.Println("Error: Duplicated values invalid solution") // if the board has duplicates value then solution invalid
			}
		}
	}
}

// input validity check
func ValidateInput(args []string) bool {
	// check if the input array has 9 elements
	if len(args) != 9 {
		fmt.Println("Error: You entered wrong number of arrays") // Invalid input
		return false
	}
	// check each element inside the array has 9 elemnts
	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			fmt.Println("Error: One or more arrays has less than 9 elements") //  invalid input
			return false
		}
	}
	// check invalid values
	for i := 0; i < len(args); i++ {
		for _, value := range args[i] {
			if (value < 49 || value > 57) && value != 46 {
				fmt.Println("Error: One or more arrays has a non integer element or a zero") // IInvalid input
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

// check the final board has no duplicates
func hasDuplicates(counter [100]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

// check for the final board solution
func isBoardValid(board *[9][9]rune) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [100]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [100]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [100]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

// check if the generated value fits in the board

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

// backtracking to solve all empty spots
func IsItSolved(table *[9][9]rune) bool {

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
						if IsItSolved(table) {
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
