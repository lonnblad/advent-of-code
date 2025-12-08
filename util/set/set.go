package set

type Set[T comparable] struct {
	values map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		values: make(map[T]bool),
	}
}

func (s *Set[T]) Add(value T) {
	s.values[value] = true
}

func (s *Set[T]) Contains(value T) bool {
	return s.values[value]
}

func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.values))

	for value := range s.values {
		values = append(values, value)
	}

	return values
}

func (s *Set[T]) Size() int {
	return len(s.values)
}

