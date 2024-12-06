package main

import "fmt"

var firstName, lastName = "Julian", "Amschwand"
var dayOfBirth, monthOfBirth, yearOfBirth = 12, 06, 2008
var numberOfSiblings = 2
var heightInMeters = 1.67
var zodiacSign = '\u264A'

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
