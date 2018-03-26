package main

import "fmt"

func main() {
  var s string

  if is_true(){
    s = fili()
  }else{
    s = coshi()
  }

  fmt.Println(s)
}

// You need to specify the return value type
func fili() string{
  return "Filib"
}

func coshi() string{
  return "Coshi"
}

func is_true() bool {
  return true
}

func is_false() bool {
  return false
}
