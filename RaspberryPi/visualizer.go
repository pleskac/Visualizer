package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Refresh rate is in seconds
func RandomVisualizer(c chan LED, refreshRate int64) {
	d := 60 * refreshRate * 1e7
	test := 0
	for {
		i := 0
		for i < MAX_ROW {
			j := 0
			for j < MAX_COL {
				//Add a white LED at that locations
				temp := LED{i, j, randomLEDValue(), randomLEDValue(), randomLEDValue()}
				temp.Print()
				//send light to be updated immediately
				c <- temp
				j++
			}
			i++
		}

		fmt.Println("TEST:", test)
		time.Sleep(time.Duration(d))
		test++
	}

	//should be impossible
	err := errors.New("Exiting infinite loop")
	panic(err)
}

func randomLEDValue() int {
	var temp int
	temp = rand.Int() % 256
	return temp
}

//Add different visualization patterns
