package main

import (
	"fmt"	
	"math"
	"math/rand"
	"time"
)

var c, python, java bool = true, false, true;

const Pi = 3.14;

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}


//reduce type
func add(x, y int, nome string) int {
	fmt.Println("Printando o nome: " + nome);
	return x + y
}

//multiple return
func swap(x, y string) (string, string) {
	return y, x
}

//clean return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Tour 01");
	rand.Seed(time.Now().UnixNano());
	fmt.Println("My favorite number is", rand.Intn(25));
	
	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 02");
	fmt.Printf("Now you have %g problems.", math.Sqrt(7));
	fmt.Println("");

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 03");
	fmt.Println(math.Pi)

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 04");
	fmt.Println(add(42, 13, "Samuel"));

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 05");
	a, b := swap( "world", "hello");
	fmt.Println(a, b);

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 06");
	fmt.Println(split(17));
	
	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 07");
	var i int;
	var y = "Samuel Santana";
	k := "GoLang";
	fmt.Println(i, c, python, java, y, k);

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 08");
	var aa, bb int = 3, 4;
	var f float64 = math.Sqrt(float64(aa*aa + bb*bb));
	var z uint = uint(f);
	fmt.Println(aa, bb, z);

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 09");
	v := 0.45 // change me!
	fmt.Printf("v is of type %T\n", v)

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 10");
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Printf("v is of type %T\n", v)

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 11");

	fmt.Println(Small);
	fmt.Println(needFloat(Big));

	fmt.Println(needInt(Small));
	fmt.Println(needFloat(Small));
	//fmt.Println(needInt(Big));
	/*fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))*/
}
