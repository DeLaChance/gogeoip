package geoip 

type GeoIpRepository interface {
	FindCountryByIpAddress(ipAddress string) *GeoIpRange
}