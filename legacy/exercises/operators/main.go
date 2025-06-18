package main

import (
	"fmt"
)

type duration int
type mile float64
type kilometer float64

func main() {
	const (
		distEarthToSun float64 = 149.6 * 1000000 * 1000
		speedOfLight           = 299792458
		m2km                   = 1.609
	)

	var hour duration = 45

	var timeEarthToSunInMin = (distEarthToSun / float64(speedOfLight)) / 60

	fmt.Println(timeEarthToSunInMin)
	fmt.Println(hour)

	var mileBerlinToParis mile = 655.3
	var kmBerlingToParis kilometer

	kmBerlingToParis = kilometer(mileBerlinToParis) * m2km
	fmt.Println(kmBerlingToParis)
}
