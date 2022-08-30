package safetypes

import (
	"encoding/json"
	"strings"
)

type Option[T any] struct {
	Value *T `json:"value,omitempty" bson:"value,omitempty" rethinkdb:"value,omitempty"`
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

func (o *Option[T]) UnwrapOr(or T) T {
	if o.IsNone() {
		return or
	}
	return *o.Value
}

func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.IsSome() {
		return json.Marshal(o.Value)
	}
	return []byte("{}"), nil
}

func (o *Option[T]) UnmarshalJSON(data []byte) error {
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		if strings.HasPrefix(string(data), "{}") {
			o.Value = nil
			return nil
		}
		return err
	}
	o.Value = &result
	return nil
}
