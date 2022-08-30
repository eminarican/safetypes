# safetypes
Rust like result and option implementation for golang

just a reminder, option type is ready for (un)marshalling, mongodb and rethinkdb so feel free to use it with ^^

## Examples

```go
import safe "github.com/eminarican/safetypes"
```

### Option
```go
func checkUnwrap(opt safe.Option[int]) {
    if opt.IsSome() {
        println(opt.Unwrap())
    } else {
        panic("poor option :(")
    }
}
```
```go
func checkUnwrapOr(opt safe.Option[int]) {
    println(opt.UnwrapOr(10))
}
```
```go
func retrunOption(some bool) (opt safe.Option[int]) {
    if some {
        return opt.Some(7)
    }
    return opt.None()
}
```
```go
type Test struct {
    Field safe.Option[int]
}

func jsonMarshal(t Test) {
    res := safe.AsResult(json.Marshal(s))
    if res.IsOk() {
        // if some: "Test{Field: 7}"
        // if none: "Test{Field: {}}"
        println(res.Unwrap())
    } else {
        panic(res.Error())
    }
}
```

### Result
```go
func checkUnwrap(res safe.Result[int]) {
    if res.IsOk() {
        println(res.Unwrap())
    } else {
        panic(res.Error())
    }
}
```
```go
func retrunResult(some bool) (res safe.Result[int]) {
    if some {
        return res.Ok(7)
    }
    return res.Err("some fancy error msg")
}
```

### Note
Error and None methods usable as structless but it doesn't infere types so instead of using `safetypes.None[T]()` and `safetypes.Err[T]("")` you could use them as how in examples above
