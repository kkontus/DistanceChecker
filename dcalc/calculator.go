package dcalc

import (
	"strconv"
	"fmt"
	"os"
	"DistanceChecker/entity"
	"DistanceChecker/config"
	"DistanceChecker/util"
)

type formula string

const (
	Haversine     formula = "Haversine"
	SomeOtherCalc formula = "SomeOtherCalc"
)

func CalculateDistance(from entity.GeoCoord, customers []entity.Customer, formula formula) []entity.Customer {
	var customersInRange []entity.Customer
	cInRange := 0
	for _, c := range customers {
		lat, _ := strconv.ParseFloat(c.Latitude, 64)
		lng, _ := strconv.ParseFloat(c.Longitude, 64)
		to := entity.GeoCoord{Lat: lat, Lng: lng}

		distance := getCalculationFormula(from, to, formula)

		if distance.Km <= config.WithinKm {
			c.Distance.Km = distance.Km
			c.Distance.Mi = distance.Mi
			c.Distance.NMi = distance.NMi

			customersInRange = append(customersInRange, c)
			cInRange++
		}
	}
	fmt.Printf("\n %d out of %d customers are within %.2f km (%.2f miles) range from Dublin office \n\n", cInRange, len(customers), config.WithinKm, util.KmToMi(config.WithinKm))

	return customersInRange
}

func getCalculationFormula(from entity.GeoCoord, to entity.GeoCoord, formula formula) entity.GeoDist {
	var distance entity.GeoDist
	switch formula {
	case Haversine:
		distance = Distance(from, to)
	case SomeOtherCalc:
		fmt.Printf("%s formula not implemented \n", SomeOtherCalc)
		os.Exit(1)
	default:
		fmt.Println("Formula not recognized")
		os.Exit(1)
	}
	return distance
}
