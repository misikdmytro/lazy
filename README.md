# Lazy

[![Lazy CI](https://github.com/misikdmytro/lazy/actions/workflows/go.yaml/badge.svg)](https://github.com/misikdmytro/lazy/actions/workflows/go.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/misikdmytro/lazy)](https://goreportcard.com/report/github.com/misikdmytro/lazy)

## Quick Start

Lazy package is package that implements lazy loading mechanism. More details [here](https://en.wikipedia.org/wiki/Lazy_loading).

The package exports single method:

```go
func NewLazy[T any](f lazyFunc[T]) Lazy[T]
```

that returns Lazy instance.

## Instalation

```bash
go get -u github.com/misikdmytro/lazy
```

## Code example

```go
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
```

Lazy initializer will be called only once.
