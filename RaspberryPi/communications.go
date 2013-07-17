package main

import (
	"fmt"
)

//can use the GPIO pins on the RaspberryPi to directly control LEDs, BUT...
//There aren't enough and it doesn't allow for expansion.
//therefore i'll use the pins to communicate with the arduinos. Probs

//Should be called serially
func UpdateLight(light LED) {
	fmt.Println("Updating LED", light.String())
}
