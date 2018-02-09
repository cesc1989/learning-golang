package main

import "fmt"

func main(){
  /*
    You can use Saiyan like this as well:
      goku := Saiyan{}
      goku := Saiyan{Name: "Goku"}
      goku.Power = 9000
  */
  goku := Saiyan{
    Name: "Goku",
    Power: 9000,
  }

  // Go passes arguments by value(copy of variable) not by reference(original variable)
  Super(goku)
  fmt.Println(goku.Power) // prints 9000

  gohan := &Saiyan{"Gohan", 3000}
  super_pointer(gohan)
  fmt.Println(gohan.Power)
}

type Saiyan struct {
  Name string
  Power int
}

func Super(s Saiyan){
  s.Power += 10000
}

func super_pointer(s *Saiyan){
  s.Power += 10000
}
