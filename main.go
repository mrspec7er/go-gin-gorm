package main

import (
	"example/go-gin/controllers"
	"example/go-gin/utility"
)

func main()  {
	// routerHandler();
	utility.DBConnection()
	controllers.UtilityRouterHandler();
}