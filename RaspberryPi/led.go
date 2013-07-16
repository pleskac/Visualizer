package main

import (
	"fmt"
)

const (
	MAX_ROW = 8
	MAX_COL = 4
)

//The data needed to represent a single RGB LED
type LED struct {
	red   int //0 to 256
	green int //0 to 256
	blue  int //0 to 256
}

// 0,0  0,1  0,2
// 1,0  1,1  1,2
// 2,0  2,1  2,2
// ROW, COL
//LED[ROW][COL]
//NEED A MUTEX/LOCK ON EACH ELEMENT
var GRID [][]LED
var GRID_READY = false

func PrintGrid() {
	for a, row := range GRID {
		fmt.Print("Row", a)
		for _, light := range row {
			fmt.Print(light, "  ")
		}
		fmt.Println()
	}
}
