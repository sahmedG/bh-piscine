package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("File name missing")
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
	
	} else { file, err := os.Open(args[0])
	if err != nil {
		fmt.Printf("Mistake: %v\n" , err.Error())
	}
	arr := make([]byte,20)
	file.Read(arr)
	fmt.Print(string(arr))
	}
}
