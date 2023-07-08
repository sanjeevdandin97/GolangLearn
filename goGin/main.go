package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

var goRoutes = [2]string{
	"example",
	"hello",
}

func main() {
	routerInit()
	createRoutes()
	runServer()
}

// Creates a gin router with default middleware
func routerInit() {
	router = gin.Default()
}

func createRoutes() {
	for i := 0; i < len(goRoutes); i++ {
		var message string = goRoutes[i]
		// A handler for GET request on routes[i]
		router.GET("/"+goRoutes[i], func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": message,
			}) // gin.H is a shortcut for map[string]interface{}
		})
	}
}

func runServer() {
	router.Run() // listen and serve on port 8080
}
