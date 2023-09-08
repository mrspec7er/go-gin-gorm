package controllers

import (
	"example/go-gin/models"
	"example/go-gin/utility"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoI struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	Desc string `json:"desc"`
	StoryPoint int `json:"storyPoint"`
}

func getTodos(c *gin.Context)  {
	todos := []models.Todo{}

	result := utility.DBConnection().Find(&todos);

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Cannot Find Todo Object!"});
		return
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func createTodos(c *gin.Context)  {
	var request TodoI
	err := c.BindJSON(&request)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Body Type"});
		return
	}
	todo := models.Todo{Name: request.Name, Completed: false, Desc: request.Desc, StoryPoint: request.StoryPoint}
	result := utility.DBConnection().Create(&todo);

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Cannot Create Todo Object!"});
		return
	}
	c.IndentedJSON(http.StatusCreated, todo)
}

func getTodo(c *gin.Context) {
	id := c.Param("id")
	todo := models.Todo{}
	result := utility.DBConnection().Find(&todo, id);

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Cannot Find Todo Object!"});
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func updateTodo(c *gin.Context) {
	id, status := c.GetQuery("id");

	if !status {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Parameter Id Required!"});
		return
	}
	todo := models.Todo{}
	result := utility.DBConnection().Find(&todo, id);

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Cannot Update Todo Object!"});
		return
	}
	updateResult := utility.DBConnection().Model(&todo).Update("Completed", true)

	if updateResult.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Cannot Update Todo Object!"});
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func deleteTodo(c *gin.Context) {
	id, status := c.GetQuery("id");

	if !status {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Parameter Id Required!"});
		return
	}
	todo := models.Todo{}
	result := utility.DBConnection().Delete(&todo, id);

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Cannot Delete Todo Object!"});
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func pointerExample(c *gin.Context) {
	x := 10
	y := *&x // create instance of the reference
	z := &x // pointing to reference

	y = y + 1 // mutating variable doesen't change the reference 
	*z = 7 // mutating variable change the reference

	fmt.Println(x,y,z)

	result := fmt.Sprintf("%v, %v, %v", x,y, *z) 

	c.IndentedJSON(http.StatusOK, result)
}


