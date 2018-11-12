package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println("Tour 01");
	fmt.Println(sqrt(2), sqrt(-4));

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 02");
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 03");
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	);

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 03");
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 04");
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
