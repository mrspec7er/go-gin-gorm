package middleware

import (
	"example/go-gin/models"
	"example/go-gin/utility"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(c *gin.Context){
	publicToken, err := c.Cookie("Authorization") 
	
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized);
		return
	}
	
	token, err := jwt.Parse(publicToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
		return []byte(os.Getenv("SECRET")), nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := models.User{}
		getUserResult := utility.DBConnection().Find(&user, claims["id"])
		if getUserResult.Error != nil || user.ID == 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Cannot Find User Object!"});
			return
		}
		c.Set("user", user)

			
		} else {
			c.AbortWithStatus(http.StatusUnauthorized);
			}
	
	c.Next()
}