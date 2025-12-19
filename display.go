package main

import (
    "fmt"
    "strconv"
    "strings"
)

// printTable prints people in a formatted table
func printTable(people []Person) {
    if len(people) == 0 {
        fmt.Println("No people in database.")
        return
    }
    
    // Calculate column widths
    nameWidth := len("Name")
    ageWidth := len("Age")
    emailWidth := len("Email")
    
    for _, person := range people {
        if len(person.Name) > nameWidth {
            nameWidth = len(person.Name)
        }
        ageStr := strconv.Itoa(person.Age)
        if len(ageStr) > ageWidth {
            ageWidth = len(ageStr)
        }
        if len(person.Email) > emailWidth {
            emailWidth = len(person.Email)
        }
    }
    
    // Print header
    printSeparator(nameWidth, ageWidth, emailWidth)
    fmt.Printf("| %-*s | %-*s | %-*s |\n", nameWidth, "Name", ageWidth, "Age", emailWidth, "Email")
    printSeparator(nameWidth, ageWidth, emailWidth)
    
    // Print rows
    for _, person := range people {
        fmt.Printf("| %-*s | %-*d | %-*s |\n", nameWidth, person.Name, ageWidth, person.Age, emailWidth, person.Email)
    }
    printSeparator(nameWidth, ageWidth, emailWidth)
}

// printSeparator prints a table separator line
func printSeparator(nameWidth, ageWidth, emailWidth int) {
    fmt.Printf("+-%-*s-+-%-*s-+-%-*s-+\n",
        nameWidth, strings.Repeat("-", nameWidth),
        ageWidth, strings.Repeat("-", ageWidth),
        emailWidth, strings.Repeat("-", emailWidth))
}
