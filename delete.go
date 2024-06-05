package main

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func Delete(db *gorm.DB, reference string) {
	ID, err := strconv.Atoi(reference)

	var result *gorm.DB

	if err == nil {
		result = db.Table("Todos").Delete(&Todo{ID: ID})
	} else {
		records := CheckRecordsExistance(db, reference)

		if len(records) == 0 {
			fmt.Println("Record(s) Doesn't Exist, try another <Name>")
			return
		}
		result = db.Table("Todos").Where("Name", reference).Delete(&Todo{})
	}

	if result.Error != nil {
		fmt.Println("Error: ", result.Error)
		return
	}

	message := "Item Deleted Successfully"

	if result.RowsAffected > 1 {
		message = "All Items with the Name " + reference + " have been delete"
	}

	fmt.Println(message)
}

func CheckRecordsExistance(db *gorm.DB, ref string) []Todo {
	var todos []Todo

	db.Limit(-1).Table("Todos").Find(&todos)

	return todos
}
