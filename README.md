# Lazy

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
var inc int
l := lazy.NewLazy(func() int {
    inc++
    return inc
})

val1 := l.Value() // returns 42
val2 := l.Value() // returns 42 as well
```

Lazy initializer will be called only once.
