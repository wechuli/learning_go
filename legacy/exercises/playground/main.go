package main

import "fmt"

func main() {

	// Question 6

	friends := []string{"Marry", "John", "Paul", "Diana"}
	friendsCopy := make([]string, len(friends))

	copy(friendsCopy, friends)

	friendsCopy[3] = "Changed"
	fmt.Println(friendsCopy)
	fmt.Println(friends)

	// Question 7

	friendsAppended := []string{}

	friendsAppended = append(friendsAppended, friends...)
	fmt.Println(friendsAppended)

	// Question 8
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}

	newYears := []int{}

	var n = len(years) - 3

	// newYears = append(newYears, years[:3]...)
	// newYears = append(newYears, years[n:]...)

	newYears = append(years[:3], years[n:]...)

	fmt.Println(newYears)

}
