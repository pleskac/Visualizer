package main

import (
	"fmt"
)

//GOOD IDEA: run the visualizer algorithm in separate thread, write to data structure, then the refresh algorithm can just get a read lock on the same datastructure. mutexes.

func main() {
	//setup the data structures
	setup()

	go Visualizer() //updates GRID based on algorithm

	//Using a package, startup the USB communication
	fmt.Println("Initialize Serial Communications")
}

func setup() {
	i := 0
	for i < MAX_ROW {
		j := 0
		var tempROW []LED
		for j < MAX_COL {
			//Add a black LED at that locations
			tempROW = append(tempROW, LED{0, 0, 0})
			j++
		}
		i++
		GRID = append(GRID, tempROW)
	}

	fmt.Println("GRID", GRID)
	fmt.Println("GRID should be set up")
	fmt.Println("GRID rows", len(GRID))

	//Sanity check
	PrintGrid()

	GRID_READY = true
}
