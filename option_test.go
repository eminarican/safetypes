package safetypes

import "testing"

func TestOption(t *testing.T) {
	if res := option_test_some(); res.IsSome() {
		t.Log("the result is", res.Unwrap())
	} else {
		panic("result should be some")
	}

	if res := option_test_some(); res.IsSome() {
		t.Log("the result is", res.Unwrap())
	} else {
		panic("result should be none")
	}
}

func option_test_some() Option[int] {
	return Some(3)
}

func option_test_none() (res Option[int]) {
	return res.None()
}
