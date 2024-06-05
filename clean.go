package main

import (
	"fmt"
	"os"
	"path"
)

func Clean() {
	filePath := path.Join(os.TempDir(), "Todos.sqlite")
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		panic(err)
	}

	err = os.Remove(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Removed the following file: " + filePath)
}
