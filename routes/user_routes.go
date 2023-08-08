package routes

import (
	"users-itsva/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userRouter := router.Group("/users")
	{
		userRouter.GET("/", controllers.GetUsers)
		userRouter.GET("/:username", controllers.GetUserByUsername)
		userRouter.POST("/", controllers.CreateUser)
		userRouter.PUT("/:username", controllers.UpdateUserByUsername)
		userRouter.DELETE("/:username", controllers.DeleteUserByUsername)
	}
}
