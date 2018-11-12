package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Tour 01");
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 02");

	var s []int = primes[1:4];
	fmt.Println(s);

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 03");

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	aa := names[0:2]
	bb := names[1:3]
	fmt.Println(aa, bb)

	bb[0] = "XXX"

	/*Uma slice não armazena todos os dados, apenas descreve uma seção de uma matriz subjacente.
	Alterando os elementos de uma slice modifica os elementos correspondentes da sua matriz subjacente.
	Outras slices que partilham a mesma matriz subjacente vão ver essas mudanças*/

	fmt.Println(aa, bb)
	fmt.Println(names)


	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 04");

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	
	//Espécie de hash map
	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(ss);

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 05");

	slice := []int{2, 3, 5, 7, 11, 13}

	slice = slice[1:4]
	fmt.Println(slice)

	slice = slice[:2]
	fmt.Println(slice)

	slice = slice[1:]
	fmt.Println(slice)


	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 06");
	sli := []int{2, 3, 5, 7, 11, 13}
	printSlice(sli)

	// Slice the slice to give it zero length.
	sli = sli[:0]
	printSlice(sli)

	// Extend its length.
	sli = s[:4]
	printSlice(sli)

	// Drop its first two values.
	sli = sli[2:]
	printSlice(sli)

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 07");
	makea := make([]int, 5)
	printSliceWithMake("a", makea)

	makeb := make([]int, 0, 5)
	printSliceWithMake("b", makeb)

	makec := makeb[:2]
	printSliceWithMake("c", makec)

	maked := makec[2:5]
	printSliceWithMake("d", maked)

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 08");
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[0][1] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}


	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 09");
	//Dimensionando uma nova matriz
	var sl []int
	printSlice(sl)

	// append works on nil slices.
	sl = append(sl, 0)
	printSlice(sl)

	// The slice grows as needed.
	sl = append(sl, 1)
	printSlice(sl)

	// We can add more than one element at a time.
	sl = append(sl, 2, 3, 4)
	printSlice(sl)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


func printSliceWithMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
