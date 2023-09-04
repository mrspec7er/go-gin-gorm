package controllers

import (
	"example/go-gin/models"
	"example/go-gin/utility"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type TodoI struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	Desc string `json:"desc"`
	StoryPoint int `json:"storyPoint"`
}

// var todos = []TodoI{
// 	{ID: 1, Name: "Learn Golang", Completed: false, Desc: "Learn go programing language to be senior developer", StoryPoint:  7},
// 	{ID: 2, Name: "Learn Business Fundamental", Completed: false, Desc: "Improve intuition in business", StoryPoint:  11},
// 	{ID: 3, Name: "Invest In Stock Market", Completed: false, Desc: "Learn about emotional control about money", StoryPoint:  5},
// }

func getTodos(c *gin.Context)  {
	todos := []models.Todo{}

	result := utility.DBConnection().Find(&todos);

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Cannot Create Todo Object!"});
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

// func getTodo(c *gin.Context) {
// 	id := c.Param("id")
// 	formatedId, err := strconv.Atoi(id);

// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Invalid Request Id"});
// 		return
// 	}
// 	todo, err := handleGetTodoByIdWithoutPointer(formatedId);

// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Todo Not Found"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, todo)
// }

// func updateTodo(c *gin.Context) {
// 	id, status := c.GetQuery("id");

// 	if !status {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Parameter Id Required!"});
// 		return
// 	}

// 	formatedId, err := strconv.Atoi(id);

// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Id"});
// 		return
// 	}
// 	todo, err := handleGetTodoByIdWithPointer(formatedId);

	
// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Todo Not Found"})
// 		return
// 	}
// 	todo.Completed = true;

// 	c.IndentedJSON(http.StatusOK, todo)
// }

// func deleteTodo(c *gin.Context) {
// 	id, status := c.GetQuery("id");

// 	if !status {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Parameter Id Required!"});
// 		return
// 	}

// 	formatedId, err := strconv.Atoi(id);

// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Id!"});
// 		return
// 	}

// 	todoIndex, err := handleGetTodoIndex(formatedId);
// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Todo Not Found!"});
// 		return
// 	}
	
// 	result := handleRemoveTodo(todoIndex, todos);

// 	todos = result

// 	c.IndentedJSON(http.StatusOK, todos)
// }

// func handleGetTodoByIdWithoutPointer(id int) (TodoI, error) {
	
// 	for _, todo := range todos {
// 		if todo.ID == id {
// 			return todo, nil // return reference object of todos!
// 		}
// 	}

// 	return TodoI{}, errors.New("Todo Not Found")
// }

// func handleGetTodoByIdWithPointer(id int) (*TodoI, error) {
	
// 	for i, todo := range todos {
// 		if todo.ID == id {
// 			return &todos[i], nil // & syntax make pointer directly return todos object so we can mutate the data!
// 		}
// 	}

// 	return nil, errors.New("Todo Not Found")
// }

// func handleGetTodoIndex(id int) (int, error) {
// 	todoInstance, err := handleGetTodoByIdWithoutPointer(id);

// 	if err != nil {
// 		return 99999, err
// 	}
	
// 	for i, todo := range todos {
// 		if todo.ID == todoInstance.ID {
// 			return i, nil
// 		}
// 	}
// 	return 99999, errors.New("Todo Not Found")
// }

// func handleRemoveTodo(index int, todos []TodoI) ([]TodoI) {
//     todos = append(todos[:index], todos[index+1:]...)
// 	return todos
// }

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
	// router.PUT("/todos", updateTodo);
	// router.DELETE("/todos", deleteTodo);
	// router.GET("/todos/:id", getTodo);

	router.Run();
}
