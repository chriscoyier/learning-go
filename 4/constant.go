package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

// Notes

// I forgot to `import "fmt"` at first and it auto-injected it here in VS code, which was nice.

// Questions

// Seems slightly weird I had to import `Println` but didn't have to import anything for `int64`. Just a built in thing?

// So not _everything_ is typed in Go. Numeric constants can be type-less until used?
