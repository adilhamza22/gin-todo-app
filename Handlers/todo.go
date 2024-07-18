package Handlers

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/Models"
	"net/http"
	"github.com/jinzhu/gorm"
)

func GetTodos(db *gorm.DB, c *gin.Context) {
	var todos []Models.Todo
	if err := db.Find(&todos).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todos})	db.Find(&todos)

}

func CreateTodo(db *gorm.DB, c *gin.Context) {
	var todo Models.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(g *gorm.DB, c *gin.Context) {
	var todo Models.Todo
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&todo)
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteTodo(db *gorm.DB, c *gin.Context) {
	id := c.Param("id")
	if err := db.Where("id = ?", id).Delete(Models.Todo{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}