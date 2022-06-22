package db_err

type errorCode uint

const (
	HashWrong errorCode = iota
)

// Implementation

type Result[T any] interface {
	placeholder() T
}

type Ok[T any] struct {
	Value T
}
type Err[T any] struct {
	Code errorCode
}

func (p Ok[T]) placeholder() T {
	panic("never called")
}
func (p Err[T]) placeholder() T {
	panic("never called")
}
