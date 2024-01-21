# optional

A simple optional type in golang with the same memory usage as a pointer to the value.

## Usage

```go
package main

import "github.com/getaddrinfo/optional"

func main() {
  some := optional.Some[int](42)
  none := optional.None[int]()

  if some.Has() {
    fmt.Printf("some has a value: %v\n", some.Value())
  }

  if !none.Has() {
    fmt.Printf("none has no value")

    // none.Value() would result in a panic when called, not when dereferencing it
  }
}
```

### JSON Marshalling

The Optional type can be used in the context of JSON un/marshalling. If the Optional is None, it will be marshalled as a null. Alternatively, you can add the `omitempty` tag to your JSON structure
to skip the field if it is None:

```go
package main

type Data struct {
  Field optional.Optional[int]
  AnotherField optional.Optional[string] `json:",omitempty"`
}

func main() {
  json.Marshal(Data{
    Field: optional.Some[int](42),
    AnotherField: optional.Some[string]("test")
  }) // {"Field":42,"AnotherField":"test"}

  json.Marshal(Data{
    Field: optional.None[int](),
    AnotherField: optional.None[string]()
  }) // {"Field":null}
}
```

