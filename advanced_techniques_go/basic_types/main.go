package main

import "fmt"

func main() {
	//Numerics()
	//errorplay()
	// NonNumerics()
	//Conversions()
	//UnicodePlay()
	//StringPlay()
	//TimePlay()
	// TimeZonePlay()
	//ConstantsPlay()
	// ArraySlicePlay()
	//CapLengthPlay()
	//SliceSectionsPlay()
	//ByteSlicePlay()
	//DeleteSlice()
	//Sorting()
	//PointersPlay()
	//RandomPlay()

	// CryptoRandPlay()
	//PhoneBookApp()
	ExercisePlay()

}

func errorplay() {
	err := Check(0, 10)
	if err == nil {
		fmt.Println("Check() ended normally")
	} else {
		fmt.Println(err)
	}

	err = Check(0, 0)
	if err.Error() == "this is a custom error message" {
		fmt.Println("Custom error detected!")
	}

	err = FormattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}
}
