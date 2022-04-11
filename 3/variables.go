package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	// Experiments
	f += "pie" // Tried := at first to see if that was similar
	fmt.Println(f)

	f = f + " is good" // can do it either way I guess
	fmt.Println(f)
}

// Notes

// It's funny how autosave will totally remove a line like `import "fmt"` if you save the file before you've typed a line that uses it.

// Questions

// Is there any other variable initialization keywords besides `var`? Like `const` and such?
// A: Next lesson taught this. `const` does exist.
