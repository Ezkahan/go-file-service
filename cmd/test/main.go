package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	select {
	case ch <- 1:
		fmt.Println("Sent 1 to channel")
	default:
		fmt.Println("No receiver ready, default case executed")
	}

}
