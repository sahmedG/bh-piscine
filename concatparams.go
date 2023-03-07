package main

import "fmt"

func ConcatParams(args []string) string {
	string1 := ""
	if len(args) > 0 {
		for i := 0; i <= len(args)-1; i++ {
			string1 = string1 + args[i] + "\n"
		}
	}
	return string1
}
