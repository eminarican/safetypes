package safetypes

type Option[T any] struct {
	value *T
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		value: &value,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		value: nil,
	}
}

func (o Option[T]) Some(value T) Option[T] {
	o.value = &value
	return o
}

func (o Option[T]) None() Option[T] {
	o.value = nil
	return o
}

func (o *Option[T]) IsSome() bool {
	return o.value != nil
}

func (o *Option[T]) IsNone() bool {
	return o.value == nil
}

func (o *Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("can't unwrap none value")
	}
	return *o.value
}
