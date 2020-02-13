package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	fmt.Println(c.licenceNo)
	return c.licenceNo
}

func (c car) Name() string {
	fmt.Println(c.brand)
	return c.brand
}

func main() {
	var myCar vehicle = car{licenceNo: "KCN 859T", brand: "Subaru Impreza"}

	myCar.License()
	myCar.Name()
}
