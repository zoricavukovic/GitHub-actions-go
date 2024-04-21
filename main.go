package main

import (
	"fmt"
)

func doStuff() string {
	return "I do stuff!"
}

func add(x uint, y uint) uint {
	return x
}

func main() {
	var x uint = 9  //0000 1001
	var y uint = 65 //0100 0001

	fmt.Println(add(y, x))

}
