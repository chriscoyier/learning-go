package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// Notes & Questions

// I tend to set the file up and run it often as I'm writing it out to see it's all in order. I ran this before I wrote `i = i + 1` and it was happy to run. Would that just run forever in the VS Code console? Would it eventually crash something?

// What scope is `j` in? The parent `main()` scope or the `for` scope?

// `continue` is interesting! I guess JavaScript has that too but I've literally never used it

// Looks like if you write an infinite loop, like `for { }`, VS Code warns you for any code after it as "unreachable" code. Reminds me of accidentally writing code after a function returns.
