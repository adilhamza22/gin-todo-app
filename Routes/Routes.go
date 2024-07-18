package Routes

import (
	"../Controllers"
	"github.com/gin-gonic/gin"

)

func setupRouter() *gin.Engine{
	r := gin.Default()
	v1 := r.Group("/v1"){
		v1.GET("/todo", Controllers.GetTodos)
		v1.POST("/todo", Controllers.CreateTodo)
		v1.GET("/todo/:id", Controllers.GetTodoByID)
		v1.PUT("/todo/:id", Controllers.UpdateTodo)
		v1.DELETE("/todo/:id", Controllers.DeleteTodo)

	}
}