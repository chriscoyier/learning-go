package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}

// From a Tour of Go
// package main

// import (
// 	"fmt"
// )

// func Sqrt(x float64) float64 {
//   z:= float64(1)
//   // z := float64(x)

//   for i:= 0; i < 10; i++ {
//     z -= (z*z - x) / (2*z)
//   }

//   return z
// }

// func main() {
// 	fmt.Println(Sqrt(2))
// }
