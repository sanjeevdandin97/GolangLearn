package getuserdetails

import "github.com/gin-gonic/gin"

func GetUserDetails(c *gin.Context) gin.H {
	return gin.H{
		"message": "hello",
	}
}
