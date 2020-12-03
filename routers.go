package main

import "go_gin_tutorial/controllers"

func InitializeRouters()  {
	router.GET("/", controllers.ShowMainPage)
}