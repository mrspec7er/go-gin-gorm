package controllers

import (
	"example/go-gin/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// script running before main function
func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}

func UtilityRouterHandler()  {
	router := gin.Default();
	
	router.GET("/todos", getTodos);
	router.POST("/todos", createTodos);
	router.PUT("/todos", updateTodo);
	router.DELETE("/todos", deleteTodo);
	router.GET("/todos/:id", getTodo);
	router.GET("/pointer", pointerExample);
	router.POST("/users", createUsers);
	router.GET("/users", getUsers);
	router.POST("/login", login);
	router.GET("/validate", middleware.Authentication, validate);

	router.Run();
}