package main

import (
	"net"

	"github.com/oschwald/geoip2-golang"
)

type geoParser struct {
	geoASN  *geoip2.Reader
	geoCity *geoip2.Reader
	// geoCountry *geoip2.Reader
}

type geonameData struct {
	valid             bool
	id                uint
	city              string
	country           string
	administratorArea string
}

type geoResult struct {
	GeoIP    string `json:"IP"`
	GeoValid bool   `json:"IsValid"`

	GeoIPAutonomousSystemNumber       uint   `json:"AutonomousSystemNumber"`
	GeoIPAutonomousSystemOrganization string `json:"AutonomousSystemOrganization"`

	GeoIPCity              string  `json:"City"`
	GeoIPCityGeoNameID     uint    `json:"CityGeoNameID"`
	GeoIPCountry           string  `json:"Country"`
	GeoIPLocationLatitude  float64 `json:"LocationLatitude"`
	GeoIPLocationLongitude float64 `json:"LocationLongitude"`
}

func newGeoParser(
	mmdbCityPath string,
	mmdbASNPath string,
) (*geoParser, error) {
	dbCity, err := geoip2.Open(mmdbCityPath)
	if err != nil {
		return nil, err
	}
	dbASN, err := geoip2.Open(mmdbASNPath)
	if err != nil {
		return nil, err
	}
	geoIPParser := geoParser{
		geoASN:  dbASN,
		geoCity: dbCity,
	}

	return &geoIPParser, nil
}

func (geoParser *geoParser) newResultFromIP(ipString string) geoResult {
	obj := geoResult{
		GeoValid: false,
	}

	ip := net.ParseIP(ipString)

	if ip == nil {
		return obj
	}

	obj.GeoIP = ip.String()

	recordCity, err := geoParser.geoCity.City(ip)

	if err == nil {
		obj.GeoValid = true
		obj.GeoIPCity = recordCity.City.Names["en"]
		obj.GeoIPCityGeoNameID = recordCity.City.GeoNameID
		obj.GeoIPCountry = recordCity.Country.IsoCode
		obj.GeoIPLocationLatitude = recordCity.Location.Latitude
		obj.GeoIPLocationLongitude = recordCity.Location.Longitude
	}

	recordASN, err := geoParser.geoASN.ASN(ip)
	if err == nil {
		obj.GeoValid = true
		obj.GeoIPAutonomousSystemOrganization = recordASN.AutonomousSystemOrganization
		obj.GeoIPAutonomousSystemNumber = recordASN.AutonomousSystemNumber
	}

	return obj
}
