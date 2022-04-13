package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Extra stuff learning elsewhere
	fmt.Println(split(17))
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// Naked return! Weird!
	return
}

// Notes

// What does `[]string` mean? Is it essentially a typed version of empty quotes?

// Why didn't they just make arrays do this stuff? Everything is like "arrays are limited, slices are better"
// https://go.dev/blog/slices-intro

// I accidentally called `s.append` at first because of muscle memory for "methods" on existing things

// So `make()` creates a slices, but so does `[]string{"g", "h", "i"}`, yes?
