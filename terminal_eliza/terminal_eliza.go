package main

import (
    "bufio"
    "fmt"
    "github.com/kennysong/goeliza"
    "os"
)

func main() {
    fmt.Println("Eliza: " + goeliza.ElizaHi())

    for {
        statement := getInput()
        fmt.Println("Eliza: " + goeliza.ReplyTo(statement))

        if goeliza.IsQuitStatement(statement) {
            break
        }
    }
}

func getInput() string {
    fmt.Print("You: ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    return input
}

