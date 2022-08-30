package safetypes

import (
	"testing"
)

func TestOption(t *testing.T) {
	if res := option_test_some(); res.IsSome() {
		t.Log("the result is", res.Unwrap())
	} else {
		panic("result should be some")
	}

	if res := option_test_none(); res.IsNone() {
		t.Log("the result is none")
	} else {
		panic("result should be none")
	}
}

func option_test_some() (opt Option[int]) {
	return opt.Some(7)
}

func option_test_none() (opt Option[int]) {
	return opt.None()
}
