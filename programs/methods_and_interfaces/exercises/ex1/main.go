package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f \n", m)
}

func (m money) printStr() (result string) {
	result = fmt.Sprintf("%.2f", m)
	return
}
func main() {
	var bankBalance money = 34.8595444
	bankBalance.print()

	fmt.Println(bankBalance.printStr())

}
