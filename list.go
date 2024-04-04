package main

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"gorm.io/gorm"
)
func List(db *gorm.DB, priority string) {
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