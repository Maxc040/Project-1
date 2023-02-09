//Maak een applicatie in Golang
//Welke een welkomstbericht (bijv.: Welkom bij Fonteyn Vakantieparken) naar de console schrijft.

package main

import (
	"fmt"
	"time"
)

//een pakket dat is geïmporteerd, hier zitten functies in etc.

func main() {
	currenTime := time.now()
	hour := currenTime.Hour()
	//hier koppel ik de variabelen aan de functies

	if hour >= 7 && hour < 12 {
		fmt.Println("Goedemorgen, Welkom bij Fonteyn Vakantieparken!")
	} else if hour >= 12 && hour < 18 {
		fmt.Println("Goedemiddag, Welkom bij Fonteyn Vakantieparken!")
	} else if hour >= 18 && hour < 23 {
		fmt.Println("Goedeavond, Welkom bij Fonteyn Vakantieparken!")
	} else if hour >= 18 && hour < 23 {
		fmt.Println("Sorry, de parkeerplaats is gesloten!")
	}
//hier zeg je in principe if het tussen dit en dit uur is dan print je dit.
//else if het iets anders is print je dit. dit heb ik voor alle tijden gedaan.
