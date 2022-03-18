package main

import "github.com/eminarican/safetypes"

func main() {

}

func test() (opt safetypes.Option[int]) {
	return opt.Some(3)
}
