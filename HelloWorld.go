package main

import (
	"fmt"
	"time"
)

//een pakket dat is geïmporteerd, hier zitten functies in etc.

func main() {
	allowedLicenses := []string{"GFA-21A", "HA2-W17", "K2X-4RT"}
	// hier maak ik een array aan met daarin 3 kentekens

	var licensePlate string
	fmt.Print("Enter license plate number: ")
	fmt.Scanln(&licensePlate)
	//hierbij laat ik de code vragen via print welke code

	accessGranted := false
	for _, allowedLicense := range allowedLicenses { //hier checkt de code of de de gegeven kenteken in de array zit
		if licensePlate == allowedLicense { //wel? dan krijg je toegang.
			accessGranted = true //dit is een for loop omdat de code dan gaat checken of de gegeven data in de array zit.
			break
		}
	}

	if !accessGranted {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein.") //hier zegt de code dat je geen toegang hebt om door te gaan
		return                                                             //omdat je iets verkeerds hebt ingevuld. accesgranted is dus niet waar.
	}

	currenTime := time.Now()
	hour := currenTime.Hour()
	//hier koppel ik de variabelen aan de functies

	if hour >= 7 && hour < 12 {
		fmt.Println("Goedemorgen, Welkom bij Fonteyn Vakantieparken!")
	} else if hour >= 12 && hour < 18 {
		fmt.Println("Goedemiddag, Welkom bij Fonteyn Vakantieparken!")
	} else if hour >= 18 && hour < 23 {
		fmt.Println("Goedeavond, Welkom bij Fonteyn Vakantieparken!")
	} else {
		fmt.Println("Sorry, de parkeerplaats is gesloten!")
	}
}

//hier zeg je in principe if het tussen dit en dit uur is dan print je dit.
//else if het iets anders is print je dit. dit heb ik voor alle tijden gedaan.
