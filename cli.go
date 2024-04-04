package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/fatih/color"
)


func CLI() {
	db := initialize()

	flag.String("Help", "", "List of possible commands")

	flag.String("Add", "", `Adds an item to the todo list
	<Name>: Must be a complete string (use "")
	<Description>: Must be a complete string (use "")
	<Priority>: Must be one of the following: (Urgent|High|Medium|Low), if not passed in, will assume Low

	Examples:
		Add "Hello World" "This is my first CLI" Urgent
		Add "Hello World 2" "The Second Hello World"
	`)

	flag.String("Delete", "", `Deletes an item from the todo list
	<Name>: Must exist in the list
	<ID>: Row number of the todo item

	Examples:
		Delete "Hello World"
		Delete 10
	`)

	flag.String("List", "", `List of all current items, Listed from highest to lowest priority`)
	flag.String("Priority", "", `List Items with priority provided

	Examples:
		Priority Urgent
		Priority Low
	`)

	flag.String("Clean", "", "Removes all current existing task by deleting database file. The next time you run a command a new db will be created")

	flag.Parse()

	main := strings.ToLower(flag.Arg(0))

	switch main {
	case "help":
		flag.PrintDefaults()
	case "add":
		name, description, priority := strings.Trim(flag.Arg(1), " "), strings.Trim(flag.Arg(2), " "), strings.Trim(flag.Arg(3), " ")
		if len(flag.Args()) > 4 {
			errorFmt := color.New(color.FgHiRed).SprintfFunc()
			fmt.Println(errorFmt("Arguments must be in a specific format\n"))
			AddScript()
			return
		}

		Add(db, name, description, priority)
	case "delete":
		reference := strings.Trim(flag.Arg(1), " ")
		if len(reference) == 0 {
			errorFmt := color.New(color.FgHiRed).SprintfFunc()
			fmt.Println(errorFmt("Pass value in order to delete a todo item, <ID> | <Name>"))
			return
		}

		Delete(db, reference)
	case "list":
		List(db, "")
	case "priority":
		priority := flag.Arg(1)
		List(db, priority)
	case "clean":
		Clean()
	default:
		errorFmt := color.New(color.FgHiRed).SprintfFunc()
		fmt.Println(errorFmt("Argument Not Found\n"))
		flag.PrintDefaults()
	}
}