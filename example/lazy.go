package main

import (
	"fmt"

	"github.com/misikdmytro/lazy"
)

func main() {
	var inc int
	l := lazy.NewLazy(func() int {
		inc++
		return inc
	})

	val1 := l.Value()
	val2 := l.Value()

	fmt.Println(val1) // 1
	fmt.Println(val2) // 1
}
