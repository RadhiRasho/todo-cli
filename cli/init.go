package main

import (
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)
func NonExistant(path string) bool {
	_, err := os.Stat(path)

	return os.IsNotExist(err)
}

func initialize() *gorm.DB {
		filePath := path.Join(os.TempDir(), "Todos.sqlite")
		exists := NonExistant(filePath)
		if exists {
			TempFile, err := os.Create(filePath)
			if err != nil {
				log.Fatal(err)
			}
			TempFile.Close()
		}

		db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{
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

		db.Exec("CREATE TABLE IF NOT EXISTS Todos (ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, Name VARCHAR(50), Description TEXT, Priority VARCHAR(6), CreatedDate DATE);")


	return db
}