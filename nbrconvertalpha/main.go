package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	mapper := func(argumet string) string {
		if os.Args[1] == "--upper" {
			switch argumet {
			case "1":
				return "A"
			case "2":
				return "B"
			case "3":
				return "C"
			case "4":
				return "D"
			case "5":
				return "E"
			case "6":
				return "F"
			case "7":
				return "G"
			case "8":
				return "H"
			case "9":
				return "I"
			case "10":
				return "J"
			case "11":
				return "K"
			case "12":
				return "L"
			case "13":
				return "M"
			case "14":
				return "N"
			case "15":
				return "O"
			case "16":
				return "P"
			case "17":
				return "Q"
			case "18":
				return "R"
			case "19":
				return "S"
			case "20":
				return "T"
			case "21":
				return "U"
			case "22":
				return "V"
			case "23":
				return "W"
			case "24":
				return "X"
			case "25":
				return "Y"
			case "26":
				return "Z"
			case "56":
				return " "
			}
		} else {
			switch argumet {
			case "1":
				return "a"
			case "2":
				return "b"
			case "3":
				return "c"
			case "4":
				return "d"
			case "5":
				return "e"
			case "6":
				return "f"
			case "7":
				return "g"
			case "8":
				return "h"
			case "9":
				return "i"
			case "10":
				return "j"
			case "11":
				return "k"
			case "12":
				return "l"
			case "13":
				return "m"
			case "14":
				return "n"
			case "15":
				return "o"
			case "16":
				return "p"
			case "17":
				return "q"
			case "18":
				return "r"
			case "19":
				return "s"
			case "20":
				return "t"
			case "21":
				return "u"
			case "22":
				return "v"
			case "23":
				return "w"
			case "24":
				return "x"
			case "25":
				return "y"
			case "26":
				return "z"
			case "56":
				return " "
			}
		}
		return " "
	}
	if os.Args[1] == "--upper" {
		for i := 2; i < len(os.Args); i++ {
			str_rune := []rune(os.Expand("${"+os.Args[i]+"}", mapper))
			z01.PrintRune(str_rune[i])
		}
	} else {
		for i := 1; i < len(os.Args); i++ {
			str_rune := []rune(os.Expand("${"+os.Args[i]+"}", mapper))
			z01.PrintRune(str_rune[i])
		}
	}

}
