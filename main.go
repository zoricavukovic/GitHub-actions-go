package main

import (
	"fmt"
)

func doStuff() string {
	return "I do stuff!"

}

func add(x string, y uint) uint {
	return y
}

func main() {
	var x uint = 9  //0000 1001
	var y uint = 65 //0100 0001

	fmt.Println(add("asda", x))
	fmt.Println(add("asda", y))
	fmt.Println("asda")
}
