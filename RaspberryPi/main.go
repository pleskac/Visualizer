package main

import (
	"fmt"
)

const (
	MAX_ROW = 8
	MAX_COL = 4
)

//GOOD IDEA: run the visualizer algorithm in separate thread, write to data structure, then the refresh algorithm can just get a read lock on the same datastructure. mutexes.

func main() {
	//setup the data structures
	setup()

	//Using a package, startup the USB communication
	fmt.Println("Initialize Serial Communications")

	fmt.Println("Visualize and refresh loop")
	go Visualizer() //updates GRID based on algorithm
	go Refresh()    //reads GRID and sends over USB
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

	//Sanity check
	for a, row := range GRID {
		fmt.Print("Row", a)
		for _, light := range row {
			fmt.Print(light, "  ")
		}
		fmt.Println()
	}
}
