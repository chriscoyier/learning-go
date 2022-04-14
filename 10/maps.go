package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// Playing... why doesn't this work?
	// myMap := make(map[int]int){1: 100, 2: 200}
	// fmt.Println("map:", myMap)
}

// Notes

// Why do you have to `make()` a map? Why not just `m := map[string]int`?

// So is a map like an Object in JavaScript? Except JavaScript has maps too so maybe it's just a map.

// The `_` there... isn't usable as a variable? Is that some special thing?

// Looks like `map[string]int` is saying "the keys are strings" and "the values are ints"
