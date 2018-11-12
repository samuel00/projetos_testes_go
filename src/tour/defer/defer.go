package main

import "fmt"

func main() {

	fmt.Println("Tour 01");
	defer fmt.Println("world")
	fmt.Println("hello")


	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 02");
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

