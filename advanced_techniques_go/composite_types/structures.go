package main

import "fmt"

type Entry struct {
	Name    string
	Surname string
	Year    int
}

func StructPlay() {
	fmt.Println(zeroS())
	fmt.Println(initS("Paul", "Wechuli", 78))
	fmt.Println(*initPtoS("Paul", "Wechuli", 45))
}

func zeroS() Entry {
	return Entry{}
}
func initS(N, S string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, Surname: S, Year: 2000}
	}

	return Entry{Name: N, Surname: S, Year: Y}
}

func zeroPtoS() *Entry {
	t := &Entry{}
	return t
}

func initPtoS(N, S string, Y int) *Entry {
	if len(S) == 0 {
		return &Entry{Name: N, Surname: "Unknown", Year: Y}
	}

	return &Entry{Name: N, Surname: S, Year: Y}
}
