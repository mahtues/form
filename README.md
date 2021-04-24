# form

Package `form` provides an easy way to unmarshal a `http.Request` form into a
Go struct.

## Usage

```Go
import (
    "net/http"
    "net/url"

    "github.com/mahtues/form" // <-- import the form package
)

// Only exported fields with the tag 'form' will be decoded.
type form struct {
    Name     string  `form:"name"`
    Age      int     `form:"age"`
    SomeBool bool    `form:"somebool"`
    Label    *string `form:"label"`
}

func main() {
    r := &http.Request{Method: http.MethodGet}
    r.URL, _ = url.Parse("http://localhost/?name=alice&age=25&somebool=true")

    var target form

    // use form.Unmarshal(*http.Request, interface{}) to decode the form into
    // the defined struct.
    if err := form.Unmarshal(r, &target); err != nil {
        // errors are due to incompatible types
        panic(err)
    }

    fmt.Println(target.Name)     // stdout: alice
    fmt.Println(target.Age)      // stdout: 25
    fmt.Println(target.SomeBool) // stdout: true
}
```

This is a simple usage for the supported primitive types:
- `string`
- `bool`
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`

## Custom Unmarshalling

`form.FieldUnmarshaler` is the interface implemented by types that requires a
customized field decoding.

```Go
type FieldUnmarshaler interface {
	UnmarshalFormField(string) error
}
```

Following example shows how to use this interface with a primitive types alias.

```Go
import (
    "net/http"
    "net/url"

    "github.com/mahtues/form" // <-- import the form package
)

type form struct {
    Name         string         `form:"name"`
    ReversedName reversedString `form:"name"` <-- type with custom decoding
}

// type implementing form.FieldUnmarshaler
type reversedString string

// form.FieldUnmarshaler implementation
func (rs *reversedString) UnmarshalFormField(s string) error {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	*rs = reversedString(runes)

	return nil
}

func main() {
    r := &http.Request{Method: http.MethodGet}
    r.URL, _ = url.Parse("http://localhost/?name=alice")

    var target form

    if err := form.Unmarshal(r, &target); err != nil {
        // errors are due to incompatible types
        panic(err)
    }

    fmt.Println(target.Name)         // stdout: alice
    fmt.Println(target.ReversedName) // stdout: ecila
}
```

Following example shows how to use this interface using a struct.

```Go
import (
    "net/http"
    "net/url"

    "github.com/mahtues/form" // <-- import the form package
)

type form struct {
    Name         string       `form:"name"`
    ReversedName simpleStruct `form:"name"` <-- type with custom decoding
}

// type implementing form.FieldUnmarshaler
type simpleStruct struct {
	s string
}

// form.FieldUnmarshaler implementation
func (rs *simpleStruct) UnmarshalFormField(s string) error {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	rs.s = string(runes)

	return nil
}

func main() {
    r := &http.Request{Method: http.MethodGet}
    r.URL, _ = url.Parse("http://localhost/?name=alice")

    var target form

    if err := form.Unmarshal(r, &target); err != nil {
        // errors are due to incompatible types
        panic(err)
    }

    fmt.Println(target.Name)           // stdout: alice
    fmt.Println(target.ReversedName.s) // stdout: ecila
}
```

## To Do

- support to slices
- support to maps
