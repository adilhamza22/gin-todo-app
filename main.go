package main

import (
	handlers "gin-todo-app/Handlers"
	"gin-todo-app/Models"
	"log"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//check whether db is empty or not
	fmt.Println("db is connected")

	db.AutoMigrate(&Models.Todo{})

	r := gin.Default()
	r.GET("/todos", handlers.GetTodos(db))
	r.POST("/todos", handlers.CreateTodo(db))
	r.PUT("/todos/:id", handlers.UpdateTodo(db))
	r.DELETE("/todos/:id", handlers.DeleteTodo(db))

	r.Run(":8080")
}
