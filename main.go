package main

import (
	"github.com/gin-gonic/gin"
	"github.com/devProgOussou/gin-gorm-rest/routes"
)

func main() {
	router := gin.New()
	routes.UserRoute(router)
	router.Run(":8080")
}