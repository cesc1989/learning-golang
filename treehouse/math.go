package main

import (
  "fmt"
  "math"
  //"time"
)

var myNumber = 1.23

func main() {
  roundedUp := math.Ceil(myNumber)
  roundedDown := math.Floor(myNumber)
  fmt.Println(roundedUp, roundedDown)
}
