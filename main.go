package main

import (
    "fmt"
    "os"

    "github.com/systemb4/expenses/expense/cli"
)

// Notes
// - json .cach/ file for default csv location and currency
// - Average daily spending in last two months
// - List on specific month and year
// - Remove an expense
// - Spaces include in name user input
// - Unit tests
// - Webapp front end
// - Integrade to show status of investments as well?

func main() {
    // Run http server by default or with '-c' run command line version
    // HttpServerRun()

    // There's a better way to handle cli args in go
    argv := os.Args[1:]

    if len(argv) <= 0 {
        fmt.Println("You must pass an argument!")
         cli.Help()

        return
    }

    err := cli.ArgsHandle(argv)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
}
