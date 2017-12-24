package data

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"DistanceChecker/entity"
)

func LoadCustomerJson(path string) []entity.Customer {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []entity.Customer
	json.Unmarshal(raw, &c)
	return c
}
