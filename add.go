package main

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

func Add(db *gorm.DB, Name, Description, Priority string) {
	if Name == "" {
		fmt.Println("Name field must contain a value")
		return
	}

	if Description == "" {
		fmt.Println("Description must contain a value")
		return
	}

	Priority = PriorityAssessment(Priority)

	todo := Todo{
		ID:          0,
		Name:        Name,
		Description: Description,
		Priority:    Priority,
		CreatedDate: time.Now().Local().Format("2006-01-02"),
	}

	result := db.Table("Todos").Omit("ID").Create(&todo)

	fmt.Println(todo.ID)
	fmt.Println(todo.Name)
	fmt.Println(todo.Description)
	fmt.Println(todo.Priority)

	if result.Error != nil {
		fmt.Println("Error: ", result.Error)
		return
	}

	fmt.Println("Item Added Successfully")
}

func PriorityAssessment(priority string) string {
	toLower := strings.ToLower(priority)
	switch toLower {
	case "urgent":
	case "high":
	case "medium":
	case "low":
		break
	default:
		toLower = "low"
	}
	return toLower
}

func AddScript() {
	fmt.Println("Adds an item to the todo list")
	fmt.Println(`
<Name>: Must be a complete string (use "")
<Description>: Must be a complete string (use "")
<Priority>: Must be one of the following: (Urgent|High|Medium|Low), if not passed in, will assume Low

Examples:
	Add "Hello World" "This is my first CLI" Urgent
	Add "Hello World 2" "The Second Hello World"
	`)
}