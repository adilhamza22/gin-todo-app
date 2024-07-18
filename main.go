package main

import (
	"fmt"
	"gin-todo-app/Models"
	"gin-todo-app/Handlers"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
	"log"

)

func main (){
	db, err := gorm.Open("mysql", "root:password@/gin-todo-app?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.Todo{})

	r:= gin.Default()
	r.GET("/todos", handlers.GetTodos)
	r.POST("/todos", handlers.CreateTodo)
	r.GET("/todos/:id", handlers.GetTodo)
	r.PUT("/todos/:id", handlers.UpdateTodo)
	r.DELETE("/todos/:id", handlers.DeleteTodo)
	r.Run(":8080")
	
}