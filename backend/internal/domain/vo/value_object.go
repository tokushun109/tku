package vo

type ValueObject[T any] interface {
	Value() T
	String() string
}
