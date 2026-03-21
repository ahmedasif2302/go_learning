package main

import "fmt"

func goMapsExample() {
	// create an empty map using map and make
	var carDetail map[string]string

	person := make(map[string]string)

	fmt.Println(carDetail == nil)
	fmt.Println(person == nil)

	person["age"] = "10"
	person["name"] = "Ahmed Asif"
	person["address"] = "Chennai"

	delete(person, "address")

	// used to check if the key exists or not ok will return bool value
	val, ok := person["address1"]
	if val == "" {
		fmt.Println("true")
	}
	fmt.Println(val, ok)

	for k, v := range person {
		fmt.Println(k, v)
	}

}

type AddressWithLatLng struct {
	building               int
	street, state, zipcode string
	lat, lng               float64
}

func goMapsWithStructs() {

	companyAddress := make(map[string]AddressWithLatLng)

	companyAddress["Google"] = AddressWithLatLng{
		building: 100,
		street:   "Newcastle",
		state:    "Brooklyn",
		zipcode:  "1003",
		lat:      90.3002,
		lng:      20.300,
	}

	companyAddress["Amazon"] = AddressWithLatLng{
		building: 200,
		street:   "Main St",
		state:    "California",
		zipcode:  "9001",
		lat:      37.3382,
		lng:      121.8863,
	}

	// before mutate
	fmt.Println(companyAddress["Google"])

	// Update Google's street from Newcastle to Seattle
	addr := companyAddress["Google"]
	addr.street = "Seattle"
	addr.lat = 102.30
	companyAddress["Google"] = addr

	// after mutate
	fmt.Println(companyAddress["Google"])

	fmt.Println(companyAddress)

}

func goMapsMain() {
	// goMapsExample()
	goMapsWithStructs()
}
