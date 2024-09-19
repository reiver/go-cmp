package cmp

func Equal[T Comparer[T]](val1 Comparer[T], val2 T) bool {
	return 0 == val1.Cmp(val2)
}
