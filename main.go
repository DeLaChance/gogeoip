package main

import (
	"go.uber.org/fx"	
	"config"
	"geoip"
)

func main() {
	fx.New(
		fx.Provide(config.InitDatabaseConnection),
		fx.Provide(geoip.InitMySqlGeoIpRepository),
		fx.Invoke(config.InitRestController),
	).Run()
}