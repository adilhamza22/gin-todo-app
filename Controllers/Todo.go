package Controllers
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"../Models"
)

func getTodos(c *gin.Context){
	var todo []Models.Todo
	err := Models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}