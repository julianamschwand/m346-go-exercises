package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	DayOfBirth   int
	MonthOfBirth int
	YearOfBirth  int
}

type Profile struct {
	NumberOfSiblings byte
	ZodiacSign       rune
	FullName         FullName
	BirthDate        BirthDate
}

func main() {
	var me = Profile{
		NumberOfSiblings: 2,
		ZodiacSign:       '\u264A',
		FullName: FullName{
			FirstName: "Julian",
			LastName:  "Amschwand",
		},
		BirthDate: BirthDate{
			DayOfBirth:   12,
			MonthOfBirth: 6,
			YearOfBirth:  2008,
		},
	}

	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)

	me.NumberOfSiblings += 1

	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
