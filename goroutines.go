package goroutines

import (
	"sync"
)

func Goroutine(wg *sync.WaitGroup, job func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		job()
	}()
}

func Launch(varargs ...func()) {
	var wg sync.WaitGroup
	for _, f := range varargs {
		Goroutine(&wg, f)
	}
	wg.Wait()
}

type result[T any] struct {
	Order int
	Value T
}

func goroutineWithResult[T any](wg *sync.WaitGroup, order int, in chan<- result[T], job func() T) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		in <- result[T]{order, job()}
	}()
}

func Async[T any](functions ...func() T) []T {
	var wg sync.WaitGroup
	stream := make(chan result[T], len(functions))
	for i, f := range functions {
		goroutineWithResult(&wg, i, stream, f)
	}

	wg.Wait()
	close(stream)

	results := make([]T, len(functions))
	for item := range stream {
		results[item.Order] = item.Value
	}
	return results
}

func Async1[T any](f1 func() T) T {
	return Async(f1)[0]
}

func Async2[T any](f1 func() T, f2 func() T) (T, T) {
	r := Async(f1, f2)
	return r[0], r[1]
}

func Async3[T any](f1 func() T, f2 func() T, f3 func() T) (T, T, T) {
	r := Async(f1, f2, f3)
	return r[0], r[1], r[2]
}

func Async4[T any](f1 func() T, f2 func() T, f3 func() T, f4 func() T) (T, T, T, T) {
	r := Async(f1, f2, f3, f4)
	return r[0], r[1], r[2], r[3]
}

func Async5[T any](f1 func() T, f2 func() T, f3 func() T, f4 func() T, f5 func() T) (T, T, T, T, T) {
	r := Async(f1, f2, f3, f4, f5)
	return r[0], r[1], r[2], r[3], r[4]
}
