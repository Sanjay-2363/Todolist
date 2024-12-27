package main

import (
	"fmt"
	"todolist/models"
	"todolist/repository"
	router "todolist/router"
)

// Function to do DB connection, automigrate models and Routing
func main() {

	repository.DbConnection()
	// Automigrate models
	err := repository.Db.AutoMigrate(&models.TodoList{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Table created successfully")
	}
	// Router Initialization
	router.RoutersFunction()
}
