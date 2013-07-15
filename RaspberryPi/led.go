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

func Refresh() {
	//iterate through the data structure
	//get a "lock" on a single LED
	//refresh it
	//needs to be an infinite loop
	for {
		for r, row := range GRID {
			for c, col := range row {
				fmt.Println("Updating cell", r, c)
				fmt.Println(col)
			}
		}
	}
}
