package config 

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"geoip"
)

func StartRestController() {
	host := "0.0.0.0" // TODO: make this configurable
	port := "8080"

	address := host + ":" + port 

	fmt.Println("Setting up rest server at %s", address)
	ginRouter := gin.Default()

	geoip.InitGeoIpRoutes(ginRouter)

	ginRouter.Run(address)
}