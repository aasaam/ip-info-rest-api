package main

import (
	"testing"
)

func TestGeo1(t *testing.T) {
	gp, gpErr := newGeoParser("/tmp/GeoLite2-City.mmdb", "/tmp/GeoLite2-ASN.mmdb")

	if gpErr != nil {
		t.Error(gpErr)
	}

	cc := gp.newResultFromIP("1.1.1.1")

	if cc.GeoIPCountry != "US" {
		t.Errorf("invalid data")
	}
}
