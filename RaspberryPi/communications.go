package main

//RIGHT NOW SERIAL PACKAGE ONLY WORKS FOR MAC OS X (I think)
import (
	"fmt"
	"bytes"
	"encoding/binary"
    "github.com/pleskac/go-serial/serial"
    _ "io"
)

// Set up options. Could be put in config file
var options serial.OpenOptions = serial.OpenOptions{
      PortName: "/dev/cu.usbmodemfd111",
      BaudRate: 19200,
      DataBits: 8,
      StopBits: 1,
      MinimumReadSize: 4,
    }


//can use the GPIO pins on the RaspberryPi to directly control LEDs, BUT...
//There aren't enough and it doesn't allow for expansion.
//therefore i'll use the pins to communicate with the arduinos. Probs

//Should be called serially
func UpdateLight(light LED) {
	fmt.Println("Updating", light.String())
	
	//Serialize the lights into bytes so we can transfer them
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, &light)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes())

    // Open the port.
    port, err := serial.Open(options)
    if err != nil {
      fmt.Println("serial.Open: %v", err)
    }

    // Make sure to close it later.
    defer port.Close()

    if(port == nil){
    	panic("Port is nil. This CAN'T work")
    }

    // Write byte array to the port.
    n, err := port.Write(buf.Bytes())
    if err != nil {
      fmt.Println("port.Write: %v", err)
    }

    fmt.Println("Wrote", n, "bytes.")
}