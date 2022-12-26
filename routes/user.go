package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/devProgOussou/gin-gorm-rest/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
}