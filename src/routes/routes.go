package routes

import (
	"github.com/gin-gonic/gin"
	"todoList/src/controller"
)

// Routes function to serve endpoints
func Routes() {
	router := gin.Default()

	router.POST("/todo", controller.CreateTodo)
	router.GET("/todo", controller.GetAllTodos)
	router.PUT("/todo/:idTodo", controller.UpdateTodo)
	router.DELETE("/todo/:idTodo", controller.DeleteTodo)

	// Run route whenever triggered
	router.Run(":8765")
}
