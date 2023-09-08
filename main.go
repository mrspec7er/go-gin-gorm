package main

import (
	"example/go-gin/controllers"
	"example/go-gin/utility"
)

func main()  {
	utility.DBConnection()
	controllers.UtilityRouterHandler();
}