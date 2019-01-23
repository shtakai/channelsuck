package main

import (
	"fmt"
)

func hell(done chan bool)  {
	fmt.Println("suckem")
	done <- true
}

func main() {
	done := make(chan bool)
	go hell(done)
	x := <- done
	//var x int
	fmt.Printf("suck ed: %v\n", x)
}
