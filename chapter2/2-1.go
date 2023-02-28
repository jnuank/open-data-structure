package chapter2

type ArrayStack[T any] struct {
	a []T
	n int
}

func New[T any]() ArrayStack[T] {
	return ArrayStack[T]{}
}

func (a ArrayStack[T]) Size() int {
	return a.n
}

func (a ArrayStack[T]) Capacity() int {
	return len(a.a)
}

func (a ArrayStack[T]) Get(i int) T {
	return a.a[i]
}

func (a *ArrayStack[T]) Set(i int, x T) {
	a.a[i] = x
}

func (a *ArrayStack[T]) Add(i int, x T) {
	if a.n+1 > len(a.a) {
		a.resize()
	}
	for j := a.n; j > i; j-- {
		a.a[j] = a.a[j-1]
	}
	a.a[i] = x
	a.n++
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func (a *ArrayStack[T]) resize() {
	newA := make([]T, max(a.n*2, 1))

	for i := 0; i < a.n; i++ {
		newA[i] = a.a[i]
	}
	a.a = newA
}

//func main() {
//	array := ArrayStack[string]{}
//	array.Add(0, "alice")
//	fmt.Printf("array[0]: %v\n", array.Get(0))
//	fmt.Printf("array capacity: %v\n", array.Capacity())
//	array.Add(1, "bob")
//	fmt.Printf("array[1]: %v\n", array.Get(1))
//	fmt.Printf("array capacity: %v\n", array.Capacity())
//	array.Add(1, "catchy")
//	fmt.Printf("array[1]: %v\n", array.Get(2))
//	fmt.Printf("array capacity: %v\n", array.Capacity())
//
//}
