package main

// findPersonIndex returns the index of person with given name, or -1 if not found
func findPersonIndex(people []Person, name string) int {
    for i, person := range people {
        if person.Name == name {
            return i
        }
    }
    return -1
}

// removePerson removes a person at given index from the slice
func removePerson(people []Person, index int) []Person {
    return append(people[:index], people[index+1:]...)
}
