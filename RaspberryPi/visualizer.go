package main

import (
	"fmt"
)

func Visualizer() {
	fmt.Println("Inside the visualizer")
	//Infinite loop to keep updating!
	for {
		fmt.Println(GRID_READY)
		if GRID_READY {
			for _, row := range GRID {
				for _, light := range row {
					r, g, b := random()
					//GET THE MUTEX
					light.red = r
					light.green = g
					light.blue = b
					//RELEASE THE MUTEX
				}
			}
			PrintGrid()
		}
		PrintGrid()
	}
}

func random() (int, int, int) {

	return 5, 5, 5
}

func Refresh() {
	//instead of Pull, lets go with PUSH
	//useless function
}

//Add different visualization patterns
