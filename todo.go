package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/glebarez/sqlite"
	"github.com/rodaine/table"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Todo struct {
	ID          int
	Name        string
	Description string
	Priority    string
	CreatedDate string
}

//go:embed db.sqlite
var dbData embed.FS

func initialize() *gorm.DB {
	var tempFile *os.File
	var err error
	if os.Getenv("DBFILE") == "" {
		tempFile, err = os.CreateTemp("", "db.*")
		if err != nil {
			log.Fatal(err)
		}
		defer tempFile.Close()

		os.Setenv("DBFILE", tempFile.Name())

		fileData, err := dbData.ReadFile("db.sqlite")
		if err != nil {
			log.Fatal(err)
		}

		if _, err := tempFile.Write(fileData); err != nil {
			log.Fatal(err)
		}
	} else {
		tempFile, err = os.Open(os.Getenv("DBFILE"))

		if err != nil {
			log.Fatal(err)
		}
	}

	db, err := gorm.Open(sqlite.Open(tempFile.Name()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		},
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		PrepareStmt:    true,
		TranslateError: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Stat(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	size := file.Size()
	if size == 0 {
		db.Exec("CREATE TABLE IF NOT EXISTS Todos (ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, Name VARCHAR(50), Description TEXT, Priority VARCHAR(6), CreatedDate DATE);")
	}


	return db

}

var db *gorm.DB

func List(priority string) {
	Todos := db.Order(`
	CASE Priority
    	WHEN 'urgent' THEN 1
    	WHEN 'high' THEN 2
    	WHEN 'medium' THEN 3
    	WHEN 'normal' THEN 4
    	WHEN 'low' THEN 5
    	ELSE 6
    END`).Table("Todos")

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Name", "Description", "Priority", "CreatedDate")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(3)

	var results []Todo
	if len(priority) > 0 {
		Todos.Limit(10).Where(&Todo{
			Priority: priority,
		}).Find(&results)
	} else {
		Todos.Limit(10).Find(&results)
	}

	for _, res := range results {
		tbl.AddRow(res.ID, res.Name, res.Description, res.Priority, res.CreatedDate)
	}

	tbl.Print()
}

func Add(Name, Description, Priority string) {
	Name = strings.Trim(Name, " ")
	if Name == "" {
		fmt.Println("Name field must contain a value")
		return
	}

	Description = strings.Trim(Description, " ")

	if Description == "" {
		fmt.Println("Description must contain a value")
		return
	}

	Priority = PriorityAssessment(Priority)

	todo := Todo{
		ID: 		0,
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

func main() {
	if db == nil {
		db = initialize()
	}

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

	flag.Parse()

	main := strings.ToLower(flag.Arg(0))

	switch main {
	case "help":
		flag.PrintDefaults()
	case "add":
		name, description, priority := flag.Arg(1), flag.Arg(2), flag.Arg(3)
		if len(flag.Args()) > 4 {
			errorFmt := color.New(color.FgHiRed).SprintfFunc()
			fmt.Println(errorFmt("Arguments must be in a specific format\n"))
			AddScript()
			return
		}

		Add(name, description, priority)
	case "list":
		List("")
	case "priority":
		priority := flag.Arg(1)
		List(priority)
	default:
		errorFmt := color.New(color.FgHiRed).SprintfFunc()
		fmt.Println(errorFmt("Argument Not Found\n"))
		flag.PrintDefaults()
	}
}
