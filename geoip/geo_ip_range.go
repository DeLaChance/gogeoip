package geoip

type GeoIpRange struct {
	ID uint 
	BeginIp int 
	EndIp int 
	Country *Country
}