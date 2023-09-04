package main

import (
	"example/go-gin/models"
	"example/go-gin/utility"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init()  {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	DB := utility.DBConnection()

	DB.AutoMigrate(&models.Todo{})
}

func main()  {
	fmt.Println("Succesfully connect the database!")
}

