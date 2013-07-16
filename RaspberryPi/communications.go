package main

import (
	"fmt"
)

//Should be called serially
func UpdateLight(light LED) {
	fmt.Println("Updating LED", light.String())
}
