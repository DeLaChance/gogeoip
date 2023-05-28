package geoip

type GeoIpRangeResponse struct {
	Query *Query `json:"query"`
	Country *Country `json:"country"`
}

type Query struct {
	IpAddress string `json:"ipAddress"`
	IpAddressNumeric uint64 `json:"ipAddressNumeric"`
}