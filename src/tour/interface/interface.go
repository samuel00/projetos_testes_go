package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	fmt.Println(a.Abs())

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 02");
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 03");
	//Interface com valor subjacente nulo

	var w *W;
	i = w;
	describe(i)
	i.M()

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 04");

	var ii interface{} = "hello"

	s := ii.(string)
	fmt.Println(s)

	s, ok := ii.(string)
	fmt.Println(s, ok)

	ff, ok := ii.(float64)
	fmt.Println(ff, ok)

	/*ff = ii.(float64) // panic
	fmt.Println(ff)*/

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 05");
	do(21)
	do("hello")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
}

}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}


type I interface {
	M()
}

type T struct {
	S string
}

type W struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

func (w *W) M() {
	if w == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(w.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}




type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
