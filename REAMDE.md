# Person Information Management System

A lightweight command-line application written in Go for managing personal contact information with persistent JSON storage.

## Overview

This project demonstrates clean Go programming practices through a simple yet functional information management system. It provides a CLI interface for adding, deleting, listing, and updating person records with data persistence using JSON file storage.

## Features

- ✅ **CRUD Operations**: Add, delete, list, and update person records
- 💾 **Persistent Storage**: JSON-based file storage for data persistence
- 🔍 **Smart Search**: List all records or search by name
- 🛡️ **Error Handling**: Comprehensive error handling throughout the application
- 📝 **Interactive Updates**: Confirmation prompts before updating existing records
- 📊 **Formatted Output**: Clean table-based display for easy data reading
- 🏗️ **Modular Design**: Separation of concerns with distinct packages for different functionality

## Project Structure
.
  ├── main.go # Entry point and command routing
  ├── commands.go # Command handlers (add, del, list)
  ├── storage.go # JSON file I/O operations
  ├── display.go # Table formatting and output display
  ├── person.go # Helper functions for person operations
  ├── models.go # Data structure definitions
  ├── people.json # Data storage file (auto-generated)
  └── README.md # This file


## Tech Stack

- **Language**: Go 1.18+
- **Storage**: JSON file-based persistence
- **CLI**: Native Go command-line argument parsing
- **I/O**: Standard library (`os`, `encoding/json`, `bufio`)

## Installation

### Prerequisites
- Go 1.18 or higher installed on your system

### Clone and Build
Clone the repository
git clone https://github.com/yehao622/person-info-system.git
cd person-info-system

Run directly
go run . <command> [arguments]

Or build executable
go build -o person-manager .
./person-manager <command> [arguments]


## Usage

### Add a Person
go run . add <name> <age> <email>

Example
go run . add Bob 20 bob@gmail.com
**Output:**
✓ Successfully added: Bob (age 20, email bob@gmail.com)
Total people in database: 1


If the person already exists, you'll be prompted to confirm the update:
Person with name 'Bob' already exists:
Current: Age=20, Email=bob@gmail.com
New: Age=25, Email=bob.smith@gmail.com
Do you want to update? (Y/y to confirm): Y
✓ Person updated successfully!


### Delete a Person
go run . del <name>

Example
go run . del Bob
**Output:**
✓ Successfully deleted: Bob
Total people in database: 0

### List All People
go run . list
**Output:**
+--------+-----+-------------------+
| Name | Age | Email |
+--------+-----+-------------------+
| Bob | 20 | bob@gmail.com |
| Alice | 25 | alice@example.com |
+--------+-----+-------------------+
Total: 2 people

### List Specific Person
go run . list <name>

Example
go run . list Bob
**Output:**
+------+-----+---------------+
| Name | Age | Email |
+------+-----+---------------+
| Bob | 20 | bob@gmail.com |
+------+-----+---------------+
Total: 1 people

## Data Model

type Person struct {
Name string json:"name"
Age int json:"age"
Email string json:"email"
}


## Key Implementation Details

### Modular Architecture
The application follows clean code principles with separation of concerns:
- **`main.go`**: Command routing and program entry point
- **`commands.go`**: Business logic for each command
- **`storage.go`**: File I/O operations and JSON marshaling/unmarshaling
- **`display.go`**: Presentation layer with formatted table output
- **`person.go`**: Helper utilities for person operations
- **`models.go`**: Data structure definitions

### Error Handling
Comprehensive error handling ensures robustness:
people, err := readPeople()
if err != nil {
return fmt.Errorf("error reading data: %v", err)
}

### Data Persistence
Records are stored in `people.json` with proper formatting:
[
{
"name": "Bob",
"age": 20,
"email": "bob@gmail.com"
},
{
"name": "Alice",
"age": 25,
"email": "alice@example.com"
}
]


### User Experience Features
- Interactive confirmation for updates to prevent accidental data overwrites
- Clear success/error messages with visual indicators (✓, ✗)
- Formatted table output for easy reading
- Helpful usage instructions on invalid input

## Code Quality Features

- ✅ Idiomatic Go code following standard conventions
- ✅ Proper error handling and propagation
- ✅ Clean separation of concerns
- ✅ Clear function naming and comments
- ✅ Efficient slice operations for data manipulation
- ✅ JSON marshaling with proper indentation
- ✅ User-friendly CLI interface

## Example Workflow

Add multiple people
go run . add Bob 20 bob@gmail.com
go run . add Alice 25 alice@example.com
go run . add Charlie 30 charlie@test.com

List all
go run . list

Search for specific person
go run . list Bob

Update person (will prompt for confirmation)
go run . add Bob 21 bob.new@gmail.com

Delete a person
go run . del Charlie

Verify deletion
go run . list


## Learning Outcomes

This project demonstrates proficiency in:
- Go syntax and standard library usage
- Command-line application development
- File I/O and JSON handling
- Error handling patterns in Go
- Clean code architecture and modularity
- User input validation and interaction
- Data structure design and manipulation

## Future Enhancements

Potential improvements for extended learning:
- [ ] Add database support (PostgreSQL, SQLite)
- [ ] Implement search with filters (age range, email domain)
- [ ] Add export functionality (CSV, XML)
- [ ] REST API wrapper
- [ ] Unit tests for all functions
- [ ] Input validation (email format, age range)
- [ ] Sorting options for list command
- [ ] Pagination for large datasets
- [ ] Configuration file support
- [ ] Logging functionality

## Contributing

This is a personal learning project, but suggestions and improvements are welcome! Feel free to:
1. Fork the repository
2. Create a feature branch
3. Submit a pull request

## License

MIT License - feel free to use this code for learning and development purposes.

## Author

**Howard (Hao) Ye**
- GitHub: [@yehao622](https://github.com/yehao622)
- Email: hyedailyuse@gmail.com
- Portfolio: [howardye.vercel.app](https://howardye.vercel.app)

## Acknowledgments

Built as a learning project to demonstrate Go programming fundamentals and clean code practices.

---

**Note**: This is a demonstration project showcasing Go programming skills. For production use, consider adding input validation, comprehensive testing, and more robust error handling.

