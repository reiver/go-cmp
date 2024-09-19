package cmp

type Comparer[T any] interface {
	Cmp(T) int
}
