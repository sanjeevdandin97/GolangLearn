package server

import (
	"fmt"
	"goGinGorm/insert"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type H map[string]any

var router *gin.Engine

// Creates a gin router with default middleware
func RouterInit(goRoutes []string, db *gorm.DB) *gin.Engine {
	router = gin.Default()
	createRoutes(goRoutes, db)
	runServer()
	return router
}

func createRoutes(goRoutes []string, db *gorm.DB) {
	for i := 0; i < len(goRoutes)-1; i++ {
		// A handler for GET request on routes[i]
		router.GET("/"+goRoutes[i], func(c *gin.Context) {
			fmt.Println()
			resp := GoRoutesCall(c.Request.URL.String(), db)
			c.JSON(200, resp) // gin.H is a shortcut for map[string]interface{}
		})
	}
}

func runServer() {
	router.Run() // listen and serve on port 8080
}

func GoRoutesCall(goRoute string, db *gorm.DB) string {
	var returnValue string
	if goRoute == `/insert` {
		returnValue = insert.InsertRecord(db)
	}
	return returnValue
}
