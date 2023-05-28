package geoip 

import (	
	"errors"
	"database/sql"
	"net"
	"math/big"
)

type MySqlGeoIpRepository struct {	
	db *sql.DB
}

func InitMySqlGeoIpRepository(db *sql.DB) *MySqlGeoIpRepository {
	return &MySqlGeoIpRepository{db: db}
}

func (repo MySqlGeoIpRepository) FindCountryByIpAddress(ipAddress string) (*GeoIpRange, error) {

	// TODO: write to file
	sql := "select c.code, c.full_name, gir.beginIp, gir.endIp from gogeoip.countries c join gogeoip.geo_ip_ranges gir on gir.countryCode = c.code where gir.beginIp <= ? and gir.endIp >= ?;"
	ipAddressNumeric := convertIpv4IpAddressToNumeric(ipAddress)

	var country Country
	var geoIpRange GeoIpRange
	geoIpRange.Country = &country

	err := repo.db.QueryRow(sql, ipAddressNumeric, ipAddressNumeric).Scan(&country.Code, &country.FullName, &geoIpRange.BeginIp, &geoIpRange.EndIp)

	if (err == nil) {
		return &geoIpRange, err
	} else {
		return &GeoIpRange{}, errors.New("An error has occurred") // TODO: better error handling.
	}
}

func convertIpv4IpAddressToNumeric(ipv4Address string) string {
	bytes := net.ParseIP(ipv4Address).To4()
	bigInt := big.NewInt(0)
	bigInt.SetBytes(bytes)
	return bigInt.String()
}
