package geoip

import (
	"github.com/gin-gonic/gin"	
	"fmt"
	"net/http"
)

func InitGeoIpRoutes(router *gin.Engine) {

	fmt.Println("Adding GEO ip rest api routes")
	router.GET("/api/countries", func(c *gin.Context) {
		c.String(http.StatusOK, "{\"result\": \"ok\"}")
	})
}