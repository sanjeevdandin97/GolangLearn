package main

import (
	"goGin/constants"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var endPoints = constants.EndPoints

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
	for _, endPoint := range endPoints {
		if endPoint.Method == "GET" {
			handleGET(endPoint)
		}
	}
}

func runServer() {
	router.Run() // listen and serve on port 8080
}

// A handler for GET request
func handleGET(endPoint constants.EndpointStruct) {
	var endPointName = "/" + endPoint.Name
	router.GET(endPointName, func(ginContext *gin.Context) {
		ginContext.JSON(200, endPoint.Response(ginContext))
	})
}
