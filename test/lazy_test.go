package test

import (
	"sync"
	"testing"

	"github.com/misikdmytro/lazy/pkg/lazy"
)

func TestLazySingleThread(t *testing.T) {
	var inc int
	l := lazy.NewLazy(func() int {
		inc++
		return inc
	})

	val1 := l.Value()
	val2 := l.Value()

	if val1 != 1 {
		t.Errorf("expected first value to be 1 but was %d", val1)
	}

	if val1 != val2 {
		t.Errorf("expected first and second value to be same, but was %d and %d", val1, val2)
	}
}

func TestLazyTwoThreads(t *testing.T) {
	testLazyNThreads(t, 2)
}

func TestLazyFiveThreads(t *testing.T) {
	testLazyNThreads(t, 5)
}

func TestLazyHundredThreads(t *testing.T) {
	testLazyNThreads(t, 100)
}

func TestLazyThousandThreads(t *testing.T) {
	testLazyNThreads(t, 1000)
}

func testLazyNThreads(t *testing.T, n int) {
	var inc int
	l := lazy.NewLazy(func() int {
		inc++
		return inc
	})

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			val := l.Value()
			if val != 1 {
				t.Errorf("expected value to be 1 but was %d", val)
			}
		}()
	}

	wg.Wait()
}
