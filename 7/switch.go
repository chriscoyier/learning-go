package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}

// Notes

// What's the difference between `Print` and `Println`? I see the output as `Write 2 as %     ` — just doesn't break the line when done?

// Weird to me that `switch` even exists at all since it seems like an easy one to punt to `if ` statements like so many other things.

// Is Go "white space dependant" language like Python? Do indents actually matter?

// This is making me wonder how programming languages know things like what time and day it is. Does it ultimately ask my operating system?

// Why doesn't `t := time.Now()` need a type? Does the type come from the function call?

// There is lots of stuff I'm not familar with here... `func(i interface{})` for example. Is `interface{}` the "type" for i now? Is that was is required for `i.(type)` to work? Why is `type` in parens?

// What is happening here? `("Don't know type %T\n", t)` - it looks like `t` is getting subbed into the string, but why is `T` uppercase then? Are there a such thing as template literals in Go where you can sprinkle the variables into the string?
