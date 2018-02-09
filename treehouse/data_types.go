package main

import (
  "fmt"
  "reflect"
  "time"
)

func main() {
  fmt.Println(reflect.TypeOf(1))
  fmt.Println(reflect.TypeOf(1.5))
  fmt.Println(reflect.TypeOf(false))
  fmt.Println(reflect.TypeOf(time.Now()))
}
