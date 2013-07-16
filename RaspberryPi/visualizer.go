package main

import (
	_ "fmt"
)

func Visualizer(c chan LED) {
	for {
		i := 0
		for i < MAX_ROW {
			j := 0
			for j < MAX_COL {
				//Add a white LED at that locations
				temp := LED{i, j, 256, 256, 256}
				temp.Print()
				//send light to be updated immediately
				c <- temp
				j++
			}
			i++
		}
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
