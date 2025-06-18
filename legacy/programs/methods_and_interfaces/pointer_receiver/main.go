package main

import "fmt"

type Car struct {
	brand string
	price int
}

func changeCar(c Car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c Car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *Car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
}

func main() {
	myCar := Car{brand: "Audi", price: 40000}
	changeCar(myCar, "Renault", 20000)
	fmt.Println(myCar)

	// method with value receiver
	myCar.changeCar1("Renault", 20000)
	fmt.Println(myCar)

	// changing with pointer
	myCar.changeCar2("Nissan Note", 25000)
	// (&myCar).changeCar2("Nissan Note", 25000)

	fmt.Println(myCar)

	var yourCar *Car
	yourCar = &myCar

	fmt.Println(*yourCar)

	// valid ways to call methods

	yourCar.changeCar2("BMW", 30000)
	fmt.Println(myCar)

	(*yourCar).changeCar2("Subaru", 21000)
	fmt.Println(myCar)
}
