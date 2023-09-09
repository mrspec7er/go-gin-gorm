package controllers

import (
	"example/go-gin/models"
	"example/go-gin/utility"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func login(c *gin.Context) {
	var request UserI
	if c.BindJSON(&request) != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Body Type"});
		return
	}

	user := models.User{}
	getUserResult := utility.DBConnection().First(&user, "email = ?", request.Email)
	if getUserResult.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Cannot Find User Object!"});
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password));

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Password Doesn't Match!"});
		return
	}	

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"exp": time.Now().Add(time.Minute).Unix(),
	})

	publicToken, err := token.SignedString([]byte(os.Getenv("SECRET")));

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()});
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", publicToken, 60 * 60 * 2, "", "", false, true)

	c.IndentedJSON(http.StatusOK, publicToken)
}

func validate(c *gin.Context) {
	user, status := c.Get("user")

	if !status {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to Get User Data!"});
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}