package main

import "fmt"

func main() {
  lad, result := do_it("yamcha")

  if result {
    fmt.Println(lad)
  } else {
    fmt.Println(lad)
  }
}

func do_it(s string) (string, bool) {
  if s == "yamcha" {
    return "Mocoro", true
  }else if s == "gohan" {
    return "Mocoro", false
  }else{
    return "Nada", false
  }
}
