This library maps Go variables to string maps.

- Nested struct keys are joined with `.`
- Slice and array value keys are 0-indexed (see example)
- It will print values of pointers
- If given a primitive of type `p` with value `v`, Flatten returns `{"p": "v"}`

## Background

  I was writing an API client for a service which expected query strings and returned XML. I wanted to define each type of request's input and output as Go structs and have [google/go-querystring](https://github.com/google/go-querystring) and [encoding/xml](https://golang.org/pkg/encoding/xml/) handle (un)marshaling.

  The `encoding/xml` package worked for my purposes, but the `google/go-querystring` did not format nested keys as the service expected. I could have sent a PR to `google/go-querystring`, but this was not feasible for reasons unrelated to that package -- nonetheless, it's a good package with a fair number of options for output format.

  Flat mapping Go structs is a fairly straightforward exercise with Go reflection and writing fewer than 100 lines-of-code was the path of least resistance. This library is slightly more general than `google/go-querystring` because it will output a string map rather than query params and will accept primitives in addition to structs.

## Example
```
type AStruct struct {
    A string
    B []string
    C CStruct
}

type CStruct struct {
    D int
    E bool
}

s := &AStruct{
    A: "hello",
    B: []string{"x", "y", "z"},
    C: CStruct{
        D: 7,
        E: true,
    },
}

o := flat.Flatten(s)

// value of `o`
//   map[string]string{
//     "A": "hello",
//     "B.0": "x",
//     "B.1": "y",
//     "B.2": "z",
//     "C.D": "7",
//     "C.E": "true",
//   }
```
