package main

import (
    "encoding/json"
    "os"
)

const filename = "people.json"

// readPeople reads all people from the JSON file
func readPeople() ([]Person, error) {
    var people []Person
    
    fileData, err := os.ReadFile(filename)
    if err != nil {
        // File doesn't exist yet, return empty slice
        if os.IsNotExist(err) {
            return people, nil
        }
        return nil, err
    }
    
    err = json.Unmarshal(fileData, &people)
    if err != nil {
        return nil, err
    }
    
    return people, nil
}

// writePeople writes all people to the JSON file
func writePeople(people []Person) error {
    data, err := json.MarshalIndent(people, "", "  ")
    if err != nil {
        return err
    }
    
    return os.WriteFile(filename, data, 0644)
}
