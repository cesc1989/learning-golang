package main

import "fmt"

func main() {
  /*
    I cannot use add() int because it is expected to receive integers
    and return an integer when I pass as arguments floating point numbers
  */
  // fmt.Println(add(1.1, 2.3))
  fmt.Println("Result of the addition is:", add(1, 2))
}

/*
  if I removed the return type I get
    too many arguments to return
    have (int)
    want ()
*/
func add(a, b int) int{
  return a + b
}
