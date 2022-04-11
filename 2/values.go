package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// Lesson learned: I accidently named this file ` values.go` at first, with a space in front it, and it broke everything.

// Questions

// 1. What does `package main` mean? Does the `main` part refer to the fact that the func is named `main`? It seems to matter, I can't make them both `blah`

// 2. What does `import "fmt" mean? Does it mean "format"? I guess it has methods in it? Is `Println` used like `console.log`?
