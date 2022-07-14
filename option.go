package safetypes

type Option[T any] struct {
	Value *T `json:"value,omitempty" bson:"value,omitempty"`
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		Value: &value,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		Value: nil,
	}
}

func (o Option[T]) Some(value T) Option[T] {
	o.Value = &value
	return o
}

func (o Option[T]) None() Option[T] {
	o.Value = nil
	return o
}

func (o *Option[T]) IsSome() bool {
	return o.Value != nil
}

func (o *Option[T]) IsNone() bool {
	return o.Value == nil
}

func (o *Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("can't unwrap none value")
	}
	return *o.Value
}
