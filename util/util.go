package util

import "math"

func DegToRad(d float64) float64 {
	return d * math.Pi / 180
}

func KmToMi(km float64) float64 {
	return km * 0.621371192
}

func KmToNMi(km float64) float64 {
	return km * 0.539956803
}
