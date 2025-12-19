package main

import (
    "fmt"
    "os"
)

func printUsage() {
    fmt.Println("Usage:")
    fmt.Println("  go run . add <name> <age> <email>")
    fmt.Println("  go run . del <name>")
    fmt.Println("  go run . list [name]")
    fmt.Println()
    fmt.Println("Examples:")
    fmt.Println("  go run . add Bob 20 bob@gmail.com")
    fmt.Println("  go run . del Bob")
    fmt.Println("  go run . list")
    fmt.Println("  go run . list Bob")
}

func main() {
    if len(os.Args) < 2 {
        printUsage()
        os.Exit(1)
    }
    
    command := os.Args[1]
    args := os.Args[2:]
    
    var err error
    
    switch command {
    case "add":
        err = handleAdd(args)
    case "del":
        err = handleDel(args)
    case "list":
        err = handleList(args)
    default:
        fmt.Printf("Error: Unknown command '%s'\n\n", command)
        printUsage()
        os.Exit(1)
    }
    
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}