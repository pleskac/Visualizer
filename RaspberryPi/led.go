package main

import (
	"fmt"
)

const (
	MAX_ROW = 8
	MAX_COL = 4
)

//The data needed to represent a single RGB LED
//Row and col saved in the struct so we can push these on a channel without worrying about order
type LED struct {
	row   int
	col   int
	red   int //0 to 256
	green int //0 to 256
	blue  int //0 to 256

}

func (light LED) String() string {
	return fmt.Sprintf("LED (%d,%d) RGB: %d/%d/%d", light.row, light.col, light.red, light.green, light.blue)
}

func (light LED) Print() {
	fmt.Printf("LED (%d,%d) RGB: %d/%d/%d\n", light.row, light.col, light.red, light.green, light.blue)
}
