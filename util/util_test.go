package util_test

import (
	. "DistanceChecker/util"
	"testing"
	"fmt"
)

func TestDegToRad(t *testing.T) {
	var deg90 = 90.0
	var res = DegToRad(deg90)

	var resFormatted = fmt.Sprintf("%.7f", res)
	var resValidationFormatted = fmt.Sprintf("%.7f", 1.57079633)

	if resFormatted == resValidationFormatted {
		t.Log("testDeg OK")
	} else {
		t.Error("testDeg FAILED")
	}
}

func TestKmToMi(t *testing.T) {
	var km = 100.0
	var res = KmToMi(km)

	var resFormatted = fmt.Sprintf("%.7f", res)
	var resValidationFormatted = fmt.Sprintf("%.7f", 62.1371192)

	if resFormatted == resValidationFormatted {
		t.Log("testMi OK")
	} else {
		t.Error("testMi FAILED")
	}
}

func TestKmToNMi(t *testing.T) {
	var km = 100.0
	var res = KmToNMi(km)

	var resFormatted = fmt.Sprintf("%.7f", res)
	var resValidationFormatted = fmt.Sprintf("%.7f", 53.9956803)

	if resFormatted == resValidationFormatted {
		t.Log("testNMi OK")
	} else {
		t.Error("testNMi FAILED")
	}
}
