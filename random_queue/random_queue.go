package random_queue

type RandomQueue[T any] struct {
	a []T
	n int
}

func (i RandomQueue[T]) Add(x int) {
}

func (i RandomQueue[T]) Size() interface{} {
	return len(i.a)
}
