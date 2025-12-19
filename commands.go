package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// handleAdd handles the 'add' command
func handleAdd(args []string) error {
    if len(args) < 3 {
        return fmt.Errorf("usage: add <name> <age> <email>")
    }
    
    name := args[0]
    ageStr := args[1]
    email := args[2]
    
    // Convert age to integer
    age, err := strconv.Atoi(ageStr)
    if err != nil {
        return fmt.Errorf("age must be a number")
    }
    
    // Read existing people
    people, err := readPeople()
    if err != nil {
        return fmt.Errorf("error reading data: %v", err)
    }
    
    // Create new person
    newPerson := Person{
        Name:  name,
        Age:   age,
        Email: email,
    }
    
    // Check if person already exists
    existingIndex := findPersonIndex(people, name)
    
    if existingIndex != -1 {
        // Person exists, prompt for update
        existing := people[existingIndex]
        fmt.Printf("Person with name '%s' already exists:\n", name)
        fmt.Printf("  Current: Age=%d, Email=%s\n", existing.Age, existing.Email)
        fmt.Printf("  New:     Age=%d, Email=%s\n", age, email)
        fmt.Print("Do you want to update? (Y/y to confirm): ")
        
        reader := bufio.NewReader(os.Stdin)
        response, err := reader.ReadString('\n')
        if err != nil {
            return fmt.Errorf("error reading input: %v", err)
        }
        
        response = strings.TrimSpace(response)
        if response == "Y" || response == "y" {
            people[existingIndex] = newPerson
            fmt.Println("✓ Person updated successfully!")
        } else {
            fmt.Println("✗ Update cancelled. No changes made.")
            return nil
        }
    } else {
        // Add new person
        people = append(people, newPerson)
        fmt.Printf("✓ Successfully added: %s (age %d, email %s)\n", name, age, email)
    }
    
    // Save to file
    if err := writePeople(people); err != nil {
        return fmt.Errorf("error saving data: %v", err)
    }
    
    fmt.Printf("Total people in database: %d\n", len(people))
    return nil
}

// handleDel handles the 'del' command
func handleDel(args []string) error {
    if len(args) < 1 {
        return fmt.Errorf("usage: del <name>")
    }
    
    name := args[0]
    
    // Read existing people
    people, err := readPeople()
    if err != nil {
        return fmt.Errorf("error reading data: %v", err)
    }
    
    // Find person
    existingIndex := findPersonIndex(people, name)
    if existingIndex == -1 {
        fmt.Printf("✗ Person with name '%s' not found. No changes made.\n", name)
        return nil
    }
    
    // Remove person
    people = removePerson(people, existingIndex)
    fmt.Printf("✓ Successfully deleted: %s\n", name)
    
    // Save to file
    if err := writePeople(people); err != nil {
        return fmt.Errorf("error saving data: %v", err)
    }
    
    fmt.Printf("Total people in database: %d\n", len(people))
    return nil
}

// handleList handles the 'list' command
func handleList(args []string) error {
    // Read existing people
    people, err := readPeople()
    if err != nil {
        return fmt.Errorf("error reading data: %v", err)
    }
    
    if len(args) == 0 {
        // List all people
        printTable(people)
    } else {
        // List specific person
        name := args[0]
        existingIndex := findPersonIndex(people, name)
        
        if existingIndex == -1 {
            fmt.Printf("User %s not find!\n", name)
        } else {
            printTable([]Person{people[existingIndex]})
        }
    }
    
    return nil
}
