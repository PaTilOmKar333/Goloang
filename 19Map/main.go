package main

import (
	"fmt"
)

const (
	permanentAddress = "permament address"
	currentAddress   = "current address"
)

type address struct {
	city string
}

type employee struct {
	firstName string
	lastName  string
	age       int
	address   map[string]address
}

func main() {

	e1 := Newemployee("omkar", "patil", 27)
	fmt.Println(e1)

	e1.address[currentAddress] = newAddress("pune")

	addrss, err := getAddress(currentAddress, e1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("address of an employee %v is %v \n", fmt.Sprintf("%v %v", e1.firstName, e1.lastName), addrss.city)

	e1.address[permanentAddress] = newAddress("goa")

	printAddress(e1)

	deleteAddress(currentAddress, e1)
}

func printAddress(e employee) {
	for typ, address := range e.address {
		fmt.Printf("%v : %v\n", typ, address.city)
	}
}

func getAddress(typ string, e employee) (addrs address, err error) {
	addrs, ok := e.address[typ]
	if !ok {
		err = fmt.Errorf("%v does not exist for employee %v", typ, fmt.Sprintf("%v %v", e.firstName, e.lastName))
		return
	}
	return
}

func deleteAddress(typ string, e employee) (addrs address, err error) {
	_, ok := e.address[typ]
	if !ok {
		err = fmt.Errorf("%v does not exist for employee %v", typ, fmt.Sprintf("%v %v", e.firstName, e.lastName))
		return
	}

	delete(e.address, typ)
	return
}

func Newemployee(fn, ln string, ag int) (emp employee) {
	emp = employee{
		firstName: fn,
		lastName:  ln,
		age:       ag,
	}
	emp.address = make(map[string]address)

	return
}

func newAddress(cityname string) address {
	return address{
		city: cityname,
	}
}
