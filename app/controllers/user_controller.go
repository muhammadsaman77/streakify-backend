package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadsaman77/streakify-backend/app/domain/dto"
	"github.com/muhammadsaman77/streakify-backend/app/services"
)

type UserController interface {
	LoginUser(c *gin.Context)	
}
type UserControllerImpl struct {
	userService services.UserService
}
func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (controller *UserControllerImpl) LoginUser(c *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, dto.ApiResponse[string]{Message: "Bad Request",Payload: err.Error()})
		return
	}
	loginResponse, err := controller.userService.LoginUser(c, &loginRequest)
	if err != nil {
		if err.Error() == "invalid credentials" {
			c.JSON(401, dto.ApiResponse[string]{Message: "Unauthorized",Payload: err.Error()})
			return
		}
		c.JSON(500, dto.ApiResponse[string]{Message: "Internal Server Error",Payload: err.Error()})
		return
	}
	c.JSON(200, dto.ApiResponse[dto.LoginResponse]{Message: "Success Login",Payload: *loginResponse})
}
