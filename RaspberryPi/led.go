package main

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
var GRID [][]LED

func Refresh() {
	//iterate through the data structure
	//get a "lock" on a single LED
	//refresh it
}
