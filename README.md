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
    // the provided stuct.
    if err := form.Unmarshal(r, &target); err != nil {
        // errors are due to incompatible types
        panic(err)
    }

    fmt.Println(target.Name)     // stdout: alice
    fmt.Println(target.Age)      // stdout: 25
    fmt.Println(target.SomeBool) // stdout: true
}
```
