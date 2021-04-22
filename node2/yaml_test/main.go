package main

import (
    "fmt"
    "strings"

    "github.com/goccy/go-yaml"
    "gopkg.in/go-playground/validator.v9"
)

type Person struct {
    Name string `validate:"required"`
    Age  int    `validate:"gte=0,lt=120"`
}

func main() {
    yml := `---
- name: john
  age: 20
- name: tom
  age: -1
- name: ken
  age: 10
`
    validate := validator.New()
    dec := yaml.NewDecoder(
        strings.NewReader(yml),
        yaml.Validator(validate),
    )
    var v []*Person
    err := dec.Decode(&v)
    fmt.Println(yaml.FormatError(err, true, true))
}
