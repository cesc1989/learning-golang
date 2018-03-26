package main

import "fmt"

func main() {
  // functions can have two return values
  lad, result := do_it("yamcha")

  // use the boolean expresion of one the returned values
  // to make the conditional look better
  if result {
    fmt.Println(lad)
  } else {
    fmt.Println(lad)
  }
}

// return values specified in return options of the function: a string and a boolean
func do_it(s string) (string, bool) {
  if s == "yamcha" {
    // return the values separated by comma
    return "Mocoro", true
  }else if s == "gohan" {
    return "Mocoro", false
  }else{
    return "Nada", false
  }
}
