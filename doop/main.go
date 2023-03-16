package main

import (
	"os"
)

func convertInttoRune(nbr int) string {
	numStr := ""
	sign := ""
	if nbr < 0 {
		nbr = nbr * -1
		sign = "-"
	}
	for nbr > 0 {
		numStr = string(rune(nbr%10)+'0') + numStr
		nbr /= 10
	}
	return sign + numStr
}

func main() {
	arguments := os.Args[1:]
	length := 0
	for i := range arguments {
		length = i + 1
	}
	if length != 3 {
		return
	}
	sign := 0
	if arguments[1] == "+" {
		sign = 0
	} else if arguments[1] == "-" {
		sign = 1
	} else if arguments[1] == "*" {
		sign = 2
	} else if arguments[1] == "/" {
		sign = 3
	} else if arguments[1] == "%" {
		sign = 4
	} else {
		return
	}
	for i, s := range arguments[0] {
		if (s >= '0' && s <= '9') || (i == 0 && s == '-') {
			continue
		} else {
			return
		}
	}
	for i, s := range arguments[2] {
		if (s >= '0' && s <= '9') || (i == 0 && s == '-') {
			continue
		} else {
			return
		}
	}
	firstNbr := Atoi(arguments[0])
	secondNbr := Atoi(arguments[2])
	if secondNbr == 0 && arguments[1] == "/" {
		os.Stderr.WriteString("No division by 0")
		os.Stderr.WriteString("\n")
		return
	}
	if secondNbr == 0 && arguments[1] == "%" {
		os.Stderr.WriteString("No Modulo by 0")
		os.Stderr.WriteString("\n")
		return
	}
	result := 0
	arrayOfFunctions := []func(int, int) int{plus, minus, times, div, mod}
	result = apply(arrayOfFunctions[sign], firstNbr, secondNbr)
	if result >= 9223372036854775807 || result <= -9223372036854775808 {
	} else {
		count := convertInttoRune(result)
		os.Stderr.WriteString(string(count))
		os.Stderr.WriteString("\n")
	}
}

func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func times(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func mod(a, b int) int {
	return a % b
}

func apply(f func(int, int) int, a int, b int) int {
	return f(a, b)
}

func Atoi(s string) int {
	runes := []rune(s)
	LenRune := 0
	result := 0
	for i := range runes {
		LenRune = i + 1
	}
	if LenRune == 0 {
		return 0
	}
	tens := 1
	for k := 0; k < LenRune-1; k++ {
		if runes[k] == '+' || runes[k] == '-' {
			continue
		}
		tens *= 10
	}
	for i := range runes {
		if (runes[0] == '-' || runes[0] == '+') && i == 0 {
			continue
		}
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		numb := 0
		for j := '0'; j < runes[i]; j++ {
			numb++
		}
		result += numb * tens
		tens /= 10
	}
	if runes[0] == '-' {
		result *= -1
	}
	return result
}
