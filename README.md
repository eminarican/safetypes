# safetypes
Rust like result and option implementation for golang

just a reminder, option type is ready for (un)marshalling, mongodb and rethinkdb so feel free to use it with ^^
```
go get github.com/eminarican/safetypes@0.0.1
```

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

### Note
Error and None methods usable as structless but it doesn't infere types so instead of using `safetypes.None[T]()` and `safetypes.Err[T]("")` you could use them as how in examples above
