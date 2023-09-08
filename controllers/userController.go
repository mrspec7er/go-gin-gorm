package controllers

import (
	"example/go-gin/models"
	"example/go-gin/utility"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserI struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Status int `json:"status"`
}

func getUsers(c *gin.Context)  {
	users := []models.User{}

	result := utility.DBConnection().Find(&users);

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Cannot Find Todo Object!"});
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func createUsers(c *gin.Context)  {
	var request UserI

	if c.BindJSON(&request) != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Body Type"});
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10);

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Body Type"});
		return
	}
	

	user := models.User{Username: request.Username, Email: request.Email, Password: string(encryptedPassword), Status: true}
	result := utility.DBConnection().Create(&user);

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Cannot Create User Object!"});
		return
	}
	c.IndentedJSON(http.StatusCreated, user)
}