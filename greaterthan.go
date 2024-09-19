package cmp

func GreaterThan[T Comparer[T]](val1 Comparer[T], val2 T) bool {
	return 1 == val1.Cmp(val2)
}
