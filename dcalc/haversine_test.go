package dcalc_test

import (
	"DistanceChecker/entity"
	"DistanceChecker/config"
	"DistanceChecker/dcalc"
	"testing"
	"fmt"
)

func TestDistance(t *testing.T) {
	officeCoord := entity.GeoCoord{Lat: config.DublinOfficeLat, Lng: config.DublinOfficeLng}
	londonCoord := entity.GeoCoord{Lat: 51.5123443, Lng: -0.09098519999997734}

	distance := dcalc.Distance(officeCoord, londonCoord)
	// there can be a slight difference in distance so I used https://www.movable-type.co.uk/scripts/latlong.html
	// and set min, max distance from Dublin office to London
	min := 463.0
	max := 465.0

	if (distance.Km >= min) && (distance.Km <= max) {
		message := fmt.Sprintf("testDistanceKm OK: Distance from Dublin to London is in acceptable range: %.2fkm", distance.Km)
		t.Log(message)
	} else {
		t.Error("testDistanceKm FAILED")
	}
}
