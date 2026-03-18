package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    
    correctPassword := "12345"

    
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter password: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input) 
 
    if input == correctPassword {
        fmt.Println("Access Granted ✅")
    } else {
        fmt.Println("Access Denied ❌")
    }
}