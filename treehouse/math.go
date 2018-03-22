package main

// You can import several packages
// every imported package has to be used otherwise
// compilation will fail
import (
  "fmt"
  "math"
)

// Not everything has to happen inside the main() entrypoint
var myNumber = 1.23

func main() {
  roundedUp := math.Ceil(myNumber)
  roundedDown := math.Floor(myNumber)
  fmt.Println(roundedUp, roundedDown)
}
