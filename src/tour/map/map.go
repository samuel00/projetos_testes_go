package main

import (
	"fmt" 
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex


var m2 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}


func main() {
	fmt.Println("Tour 01");
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m2["Google"])

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 02");
	mapInt := make(map[string]int)

	mapInt["Answer"] = 42
	fmt.Println("The value:", mapInt["Answer"])

	mapInt["Answer"] = 48
	fmt.Println("The value:", mapInt["Answer"])

	delete(mapInt, "Answer")
	fmt.Println("The value:", mapInt["Answer"])

	v, ok := mapInt["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	fmt.Println("");
	fmt.Println("#######################################");
	fmt.Println("Tour 03");
	m3 := make(map[string]int)

	m3["Answer"] = 42;
	fmt.Println("The value:", m3["Answer"]);

	m3["Answer"] = 48;
	fmt.Println("The value:", m3["Answer"]);

	//delete(m3, "Answer")
	//fmt.Println("The value:", m3["Answer"])

	v2, k := m3["Answer"];
		
	if k == true{
		fmt.Println("Valor encontrado:", v2);
	}
}
