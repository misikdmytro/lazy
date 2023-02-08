package example

import (
	"fmt"

	"github.com/misikdmytro/lazy/pkg/lazy"
)

func LazyExample() {
	var inc int
	l := lazy.NewLazy(func() int {
		inc++
		return inc
	})

	val1 := l.Value()
	val2 := l.Value()

	fmt.Println(val1) // 42
	fmt.Println(val2) // 42
}
