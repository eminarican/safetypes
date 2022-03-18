package safetypes

import "testing"

func TestResult(t *testing.T) {
	if res := result_test_ok(); res.IsOk() {
		t.Log("the result is", res.Unwrap())
	} else {
		panic("result should be ok")
	}

	if res := result_test_none(); res.IsErr() {
		t.Log("the result is:", res.Error())
	} else {
		panic("result should be err")
	}
}

func result_test_ok() (res Result[int]) {
	return res.Ok(7)
}

func result_test_none() (res Result[int]) {
	return res.Err("some fancy error message")
}
