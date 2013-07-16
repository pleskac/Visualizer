package main

import (
	_ "fmt"
)

//GOOD IDEA: run the visualizer algorithm in separate thread, write to data structure, then the refresh algorithm can just get a read lock on the same datastructure. mutexes.

func main() {

	lightsToUpdate := make(chan LED)

	//add things to channel continuously
	go RandomVisualizer(lightsToUpdate, 2)

	//take things off channel
	//must send things over USB serially, this will ensure taking things off the channel one at a time, waiting to complete, then continuing
	for {
		UpdateLight(<-lightsToUpdate)
	}

}
