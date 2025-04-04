package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadsaman77/streakify-backend/app/controllers"
	"github.com/muhammadsaman77/streakify-backend/app/helper"
	"github.com/muhammadsaman77/streakify-backend/app/repositories"
	"github.com/muhammadsaman77/streakify-backend/app/services"
	"github.com/muhammadsaman77/streakify-backend/config"
)

func main(){
	userRepository := repositories.NewUserRepository()
	db := config.InitDB()
	userService := services.NewUserService(
		userRepository,db, helper.NewPasswordHelper(), helper.NewJWTHelper(),	
	)
	userController := controllers.NewUserController(userService)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	api := router.Group("/api")
	auth := api.Group("/auth")
	{
		auth.POST("/login", userController.LoginUser)
	}
	router.Run(":8080")
}