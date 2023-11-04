package getorgdetails

import "github.com/gin-gonic/gin"

func GetOrganizationDetails(ginContext *gin.Context) gin.H {
	return gin.H{
		"message": "OrgName",
	}
}
