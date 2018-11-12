package main

import "fmt"

type Vertex struct {
	X int
	Y int
}


var (
	vertex1 = Vertex{1, 2}  // has type Vertex
	vertex2 = Vertex{X: 1}  // Y:0 is implicit
	vertex3 = Vertex{}      // X:0 and Y:0
	ponteiro  = &Vertex{1, 2} // has type *Vertex
)


func main() {
	fmt.Println("Tour 1");
	fmt.Println(Vertex{1, 2});


	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 02");
	v := Vertex{1, 2};
	fmt.Println(v.X);
	v.X = 4;
	fmt.Println(v.X);

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 03");
	vv := Vertex{1, 2};
	p := &vv;
	p.X = 1e9;
	fmt.Println(vv);

	fmt.Println(vertex1, p, vertex2, vertex3);
}

