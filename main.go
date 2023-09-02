package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoI struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	Desc string `json:"desc"`
	StoryPoint int `json:"storyPoint"`
}

var todos = []TodoI{
	{ID: 1, Name: "Learn Golang", Completed: false, Desc: "Learn go programing language to be senior developer", StoryPoint:  7},
	{ID: 2, Name: "Learn Business Fundamental", Completed: false, Desc: "Improve intuition in business", StoryPoint:  11},
	{ID: 3, Name: "Invest In Stock Market", Completed: false, Desc: "Learn about emotional control about money", StoryPoint:  5},
}

func getTodos(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, todos)
}

func createTodos(c *gin.Context)  {
	var newTodo TodoI

	err := c.BindJSON(&newTodo)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Body Type"});
		return
	}
	newTodo.ID = len(todos) + 1;
	newTodo.Completed = false;

	result := append(todos, newTodo);

	todos = result

	c.IndentedJSON(http.StatusCreated, todos)
}

func getTodo(c *gin.Context) {
	id := c.Param("id")
	formatedId, err := strconv.Atoi(id);

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Invalid Request Id"});
		return
	}
	todo, err := handleGetTodoByIdWithoutPointer(formatedId);

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Todo Not Found"})
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

	formatedId, err := strconv.Atoi(id);

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Id"});
		return
	}
	todo, err := handleGetTodoByIdWithPointer(formatedId);

	
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Todo Not Found"})
		return
	}
	todo.Completed = true;

	c.IndentedJSON(http.StatusOK, todo)
}

func deleteTodo(c *gin.Context) {
	id, status := c.GetQuery("id");

	if !status {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Parameter Id Required!"});
		return
	}

	formatedId, err := strconv.Atoi(id);

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Id!"});
		return
	}

	todoIndex, err := handleGetTodoIndex(formatedId);
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Todo Not Found!"});
		return
	}
	
	result := handleRemoveTodo(todoIndex, todos);

	todos = result

	c.IndentedJSON(http.StatusOK, todos)
}

func handleGetTodoByIdWithoutPointer(id int) (TodoI, error) {
	
	for _, todo := range todos {
		if todo.ID == id {
			return todo, nil // return reference object of todos!
		}
	}

	return TodoI{}, errors.New("Todo Not Found")
}

func handleGetTodoByIdWithPointer(id int) (*TodoI, error) {
	
	for i, todo := range todos {
		if todo.ID == id {
			return &todos[i], nil // & syntax make pointer directly return todos object so we can mutate the data!
		}
	}

	return nil, errors.New("Todo Not Found")
}

func handleGetTodoIndex(id int) (int, error) {
	todoInstance, err := handleGetTodoByIdWithoutPointer(id);

	if err != nil {
		return 99999, err
	}
	
	for i, todo := range todos {
		if todo.ID == todoInstance.ID {
			return i, nil
		}
	}
	return 99999, errors.New("Todo Not Found")
}

func handleRemoveTodo(index int, todos []TodoI) ([]TodoI) {
    todos = append(todos[:index], todos[index+1:]...)
	return todos
}

func main()  {
	router := gin.Default();
	
	router.GET("/todos", getTodos);
	router.POST("/todos", createTodos);
	router.PUT("/todos", updateTodo);
	router.DELETE("/todos", deleteTodo);
	router.GET("/todos/:id", getTodo);

	router.Run("127.0.0.1:8080");
}