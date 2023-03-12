package main

import "fmt"

func main() {
    filename := os.Ars[1:]
    file, err := os.Open(filename)
    if err != nil {
        fmt.Printf("File name is missing : v%\n", err.Error())
    }
    arr := make([]byte)
    file.Read(arr)
    fmt.Println(string(arr))
}
