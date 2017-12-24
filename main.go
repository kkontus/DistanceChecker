package main

import (
	"fmt"
	"DistanceChecker/entity"
	"DistanceChecker/config"
	"DistanceChecker/data"
	"DistanceChecker/dcalc"
	"sort"
	"DistanceChecker/sorter"
)

func main() {
	officeCoord := entity.GeoCoord{Lat: config.DublinOfficeLat, Lng: config.DublinOfficeLng}
	customersCoord := data.LoadCustomerJson(config.DataPath)

	customersInRange := dcalc.CalculateDistance(officeCoord, customersCoord, dcalc.Haversine)

	//for _, c := range customersInRange {
	//	fmt.Printf("%d | %.2f km | %.2f mi | %.2f nmi | %s \n", c.UserId, c.Distance.Km, c.Distance.Mi, c.Distance.NMi, c.Name)
	//}
	//fmt.Println()

	fmt.Println("SORTED:")
	sort.Sort(sorter.ByUserId(customersInRange))
	for _, c := range customersInRange {
		fmt.Printf("%d | %.2f km | %.2f mi | %.2f nmi | %s \n", c.UserId, c.Distance.Km, c.Distance.Mi, c.Distance.NMi, c.Name)
	}
	fmt.Println()
}
