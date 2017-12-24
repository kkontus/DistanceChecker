package main

import (
	"fmt"
	"DistanceChecker/entity"
	"DistanceChecker/config"
	"DistanceChecker/data"
	"DistanceChecker/dcalc"
)

func main() {
	officeCoord := entity.GeoCoord{Lat: config.DublinOfficeLat, Lng: config.DublinOfficeLng}
	customersCoord := data.LoadCustomerJson(config.DataPath)

	customersInRange := dcalc.CalculateDistance(officeCoord, customersCoord, dcalc.Haversine)

	for _, c := range customersInRange {
		fmt.Printf("| %.2f km | %.2f mi | %.2f nmi | %s \n", c.Distance.Km, c.Distance.Mi, c.Distance.NMi, c.Name)
	}
	fmt.Println()
}
