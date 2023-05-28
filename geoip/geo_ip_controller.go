package geoip

import (
	"github.com/gin-gonic/gin"	
	"fmt"
	"net/http"
)

func InitGeoIpRoutes(router *gin.Engine, repository *MySqlGeoIpRepository) {

	fmt.Println("Adding GEO ip rest api routes")

	router.GET("/api/geoip/:ipAddress/country", func(c *gin.Context) {
		ipAddress := c.Param("ipAddress")
		geoIpRange, err := repository.FindCountryByIpAddress(ipAddress)
		if (err == nil) {
			c.JSON(http.StatusOK, *geoIpRange.Country)
		} else {
			c.String(http.StatusInternalServerError, "Server error")
		}		
	})
}