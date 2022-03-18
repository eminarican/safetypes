# safetypes
Rust like result and option implementation for golang

## Examples

### Option
```go
func main() {
    if opt := test(true); opt.IsSome() {
        println(opt.Unwrap())
    } else {
        println("poor option :(")
    }
}

func test(some bool) (opt safetypes.Option[int]) {
    if some {
        return opt.Some(7)
    }
    return opt.None()
}
```

### Result
```go
func main() {
    if res := test(true); res.IsOk() {
        println(res.Unwrap())
    } else {
        panic(res.Error())
    }
}

func test(ok bool) (res safetypes.Result[int]) {
    if ok {
        return res.Ok(7)
    }
    return res.Err("some fancy error msg")
}
```
