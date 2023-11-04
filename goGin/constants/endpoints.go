package constants

import (
	"goGin/getorgdetails"
	"goGin/getuserdetails"

	"github.com/gin-gonic/gin"
)

type EndpointStruct struct {
	Name     string
	Method   string
	Guarded  bool
	Response func(ginContext *gin.Context) gin.H
}

var EndPoints = []EndpointStruct{
	{
		Name:     "getuserdetails",
		Method:   "GET",
		Guarded:  false,
		Response: getuserdetails.GetUserDetails,
	},
	{
		Name:     "getorgdetails",
		Method:   "GET",
		Guarded:  false,
		Response: getorgdetails.GetOrganizationDetails,
	},
}
