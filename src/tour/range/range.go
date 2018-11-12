package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	fmt.Println("Tour 01");
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 02");

	//Range ignorando o indice
	pows := make([]int, 10)
	for i := range pows {
		pows[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pows {
		fmt.Printf("%d\n", value)
	}
}

