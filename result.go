package safetypes

import "errors"

type Result[T any] struct {
	err   error
	value *T
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: &value,
		err:   nil,
	}
}

func Err[T any](err string) Result[T] {
	return Result[T]{
		value: nil,
		err:   errors.New(err),
	}
}

func AsResult[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err.Error())
	}
	return Ok(value)
}

func (r Result[T]) Ok(value T) Result[T] {
	r.value = &value
	r.err = nil
	return r
}

func (r Result[T]) Err(err string) Result[T] {
	r.value = nil
	r.err = errors.New(err)
	return r
}

func (r *Result[T]) IsOk() bool {
	return r.err == nil
}

func (r *Result[T]) IsErr() bool {
	return r.err != nil
}

func (r *Result[T]) Error() error {
	return r.err
}

func (o *Result[T]) Unwrap() T {
	if o.IsErr() {
		panic("can't unwrap err value")
	}
	return *o.value
}
