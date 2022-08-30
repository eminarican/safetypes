# safetypes
Rust like result and option implementation for golang

just a reminder, option type is ready for (un)marshalling, mongodb and rethinkdb so feel free to use it with ^^

## Examples

### Option
```go
import safe "github.com/eminarican/safetypes"

func main() {
    opt := test(true)

    // "check and unwrap" approach
    if opt.IsSome() {
        println(opt.Unwrap())
    } else {
        panic("poor option :(")
    }

    // "unwrap or" approach
    println(opt.UnwrapOr(10))
}

func test(some bool) (opt safe.Option[int]) {
    if some {
        return opt.Some(7)
    }
    return opt.None()
}
```

#### Json
```go
import safe "github.com/eminarican/safetypes"

type Test struct {
    Field safe.Option[int]
}

func main() {
    s := Test{
	    Field: test(true),
	}
    res := safe.AsResult(json.Marshal(s))
    if res.IsOk() {
        // if some: "Test{Field: 7}"
        // if none: "Test{Field: {}}"
        println(res.Unwrap())
    } else {
        panic(res.Error())
    }
}

func test(some bool) (opt safe.Option[int]) {
    if some {
        return opt.Some(7)
    }
    return opt.None()
}
```

### Result
```go
import safe "github.com/eminarican/safetypes"

func main() {
    res := test(true)

    if res.IsOk() {
        println(res.Unwrap())
    } else {
        panic(res.Error())
    }
}

func test(ok bool) (res safe.Result[int]) {
    if ok {
        return res.Ok(7)
    }
    return res.Err("some fancy error msg")
}
```

### Note
Error and None methods usable as structless but it doesn't infere types so instead of using `safetypes.None[T]()` and `safetypes.Err[T]("")` you could use them as how in examples above
