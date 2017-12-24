package dcalc

import (
	"math"
	"DistanceChecker/entity"
	"DistanceChecker/util"
)

// https://www.movable-type.co.uk/scripts/latlong.html
const earthRadiusKm = 6371

func Distance(from entity.GeoCoord, to entity.GeoCoord) entity.GeoDist {
	latFrom := util.DegToRad(from.Lat)
	lngFrom := util.DegToRad(from.Lng)
	latTo := util.DegToRad(to.Lat)
	lngTo := util.DegToRad(to.Lng)

	deltaLat := latTo - latFrom
	deltaLng := lngTo - lngFrom

	a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(latFrom)*math.Cos(latTo)*math.Pow(math.Sin(deltaLng/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	km := c * earthRadiusKm                // kilometers
	mi := c * util.KmToMi(earthRadiusKm)   // miles
	nmi := c * util.KmToNMi(earthRadiusKm) // nautical miles

	return entity.GeoDist{Km: km, Mi: mi, NMi: nmi}
}
