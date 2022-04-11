package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Notes

// What is it exactly that auto-adds needed imports? Is that the VS Code extension?

// No parens again! Looks like if I add them and save, it removes them, but if I add them and "Save with Formatting" and run the file, it's fine, so they are optional? Must be Prettier removing them?

// Looks like the standard formatting for math is not to have any space between stuff? So `7%3` or `1+1` not `7 % 3` or `1 + 1`.

// It occurs to me the semicolons are interesting in a language that looks like it otherwise avoids them?

// I heard once that in JavaScript there aren't really `if else` statements, it's just an `if` statement where the `else` part happens to be another `if` statement. Relevant here?

// This lesson notes the answer to a question I had my first day looking at Go. Are there ternary operators? Nope. Gotta us `if`
