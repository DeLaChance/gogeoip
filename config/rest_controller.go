package config 

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"geoip"
)

func InitRestController(repository *geoip.MySqlGeoIpRepository) {
	host := "0.0.0.0" // TODO: make this configurable
	port := "8080"

	address := host + ":" + port 

	fmt.Println("Setting up rest server at: ", address)
	ginRouter := gin.Default()

	geoip.InitGeoIpRoutes(ginRouter, repository)
	ginRouter.Run(address)
}