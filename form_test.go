package form

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	t.Run("primitive int", func(t *testing.T) {
		r := buildRequest("a=4")

		type testStruct struct {
			A int `form:"a"`
		}

		expected := testStruct{
			A: 4,
		}

		var actual testStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})

	t.Run("pointer to int", func(t *testing.T) {
		r := buildRequest("a=4")
		a := 4

		type TestStruct struct {
			A *int `form:"a"`
		}

		expected := TestStruct{
			A: &a,
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})

	t.Run("pointer to pointer to int", func(t *testing.T) {
		r := buildRequest("a=-4")
		a := -4
		pa := &a

		type TestStruct struct {
			A **int `form:"a"`
		}

		expected := TestStruct{
			A: &pa,
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})
}

func TestString(t *testing.T) {
	t.Run("primitive string", func(t *testing.T) {
		r := buildRequest("sometext=text")

		type TestStruct struct {
			S string `form:"sometext"`
		}

		expected := TestStruct{
			S: "text",
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})

	t.Run("pointer to string", func(t *testing.T) {
		r := buildRequest("sometext=text")
		s := "text"

		type TestStruct struct {
			S *string `form:"sometext"`
		}

		expected := TestStruct{
			S: &s,
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})

	t.Run("pointer to pointer to string", func(t *testing.T) {
		r := buildRequest("sometext=text")
		s := "text"
		ps := &s

		type TestStruct struct {
			S **string `form:"sometext"`
		}

		expected := TestStruct{
			S: &ps,
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})
}

func TestBool(t *testing.T) {
	t.Run("primitive bool", func(t *testing.T) {
		r := buildRequest("b=true")

		type TestStruct struct {
			B bool `form:"b"`
		}

		expected := TestStruct{
			B: true,
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})

	t.Run("pointer to bool", func(t *testing.T) {
		r := buildRequest("b=true")
		b := true

		type TestStruct struct {
			B *bool `form:"b"`
		}

		expected := TestStruct{
			B: &b,
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})

	t.Run("pointer to pointer to bool", func(t *testing.T) {
		r := buildRequest("b=true")
		b := true
		pb := &b

		type TestStruct struct {
			B **bool `form:"b"`
		}

		expected := TestStruct{
			B: &pb,
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})

	t.Run("bool is false", func(t *testing.T) {
		r := buildRequest("b=false&pb=0&ppb=True")
		b := false
		pb := &b

		type TestStruct struct {
			B   bool   `form:"b"`
			PB  *bool  `form:"pb"`
			PPB **bool `form:"ppb"`
		}

		expected := TestStruct{
			B:   b,
			PB:  pb,
			PPB: &pb,
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})
}

type simpleStruct struct {
	s string
}

func (rs *simpleStruct) UnmarshalFormField(s string) error {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	rs.s = string(runes)

	return nil
}

func (rs simpleStruct) String() string {
	return rs.s
}

func TestStruct(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		r := buildRequest("sometext=text")
		s := simpleStruct{"txet"}

		type TestStruct struct {
			S simpleStruct `form:"sometext"`
		}

		expected := TestStruct{
			S: s,
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})

	t.Run("pointer to struct", func(t *testing.T) {
		r := buildRequest("sometext=text")
		s := simpleStruct{"txet"}

		type TestStruct struct {
			S *simpleStruct `form:"sometext"`
		}

		expected := TestStruct{
			S: &s,
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})

	t.Run("pointer to pointer to struct", func(t *testing.T) {
		r := buildRequest("sometext=text")
		s := simpleStruct{"txet"}
		ps := &s

		type TestStruct struct {
			S **simpleStruct `form:"sometext"`
		}

		expected := TestStruct{
			S: &ps,
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})
}

type reversedString string

func (rs *reversedString) UnmarshalFormField(s string) error {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	*rs = reversedString(runes)

	return nil
}

func (rs reversedString) String() string {
	return string(rs)
}

func TestAliasString(t *testing.T) {
	t.Run("alias string", func(t *testing.T) {
		r := buildRequest("sometext=text")
		s := reversedString("txet")

		type TestStruct struct {
			S reversedString `form:"sometext"`
		}

		expected := TestStruct{
			S: s,
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})

	t.Run("pointer to alias string", func(t *testing.T) {
		r := buildRequest("sometext=text")
		s := reversedString("txet")

		type TestStruct struct {
			S *reversedString `form:"sometext"`
		}

		expected := TestStruct{
			S: &s,
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})

	t.Run("pointer to pointer to alias string", func(t *testing.T) {
		r := buildRequest("sometext=text")
		s := reversedString("txet")
		ps := &s

		type TestStruct struct {
			S **reversedString `form:"sometext"`
		}

		expected := TestStruct{
			S: &ps,
		}

		var actual TestStruct

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})
}

func TestUnnamedType(t *testing.T) {
	t.Run("alias string", func(t *testing.T) {
		r := buildRequest("sometext=text")
		s := reversedString("txet")

		expected := struct {
			S reversedString `form:"sometext"`
		}{
			S: s,
		}

		var actual struct {
			S reversedString `form:"sometext"`
		}

		err := Unmarshal(r, &actual)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v not equal to %v", actual, expected)
		}
	})
}

func buildRequest(query string) *http.Request {
	r := &http.Request{Method: http.MethodGet}
	r.URL, _ = url.Parse("http://localhost/?" + query)
	return r
}
