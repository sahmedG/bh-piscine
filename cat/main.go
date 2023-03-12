package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Hello")
		fmt.Println("Hello")
	}
	for i := 0; i < len(args); i++ {
		data, err := ioutil.ReadFile(args[i])
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
		}
		fmt.Print(string(data))
	}
}
