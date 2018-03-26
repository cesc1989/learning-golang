package main

import "fmt"

func main(){
  /*
    You can use Saiyan like this as well:
      goku := Saiyan{}
      goku := Saiyan{Name: "Goku"}
      goku.Power = 9000
  */
  // note the trailing comma. It is required for this code to compile
  goku := Saiyan{
    Name: "Goku",
    Power: 9000,
  }

  // Go passes arguments by value(copy of variable) not by reference(original variable)
  Super(goku)
  fmt.Println(goku.Power) // prints 9000, not 19000

  // You can define the values of the structure basing in the order they're defined
  gohan := &Saiyan{"Gohan", 3000}
  super_pointer(gohan)
  fmt.Println(gohan.Power) // prints 13000, not 3000
}

// Go does not have classes but you can define structures using the struct keyword
type Saiyan struct {
  Name string
  Power int
}

// then things can be declared as being of type of given structured
// in this case, variable `s` type is Saiyan
func Super(s Saiyan){
  s.Power += 10000
}

// As Go passes arguments by value and not reference, we can use Pointers to accomplish the latter
// You need to use the _address of_ operator (&) and the _pointer to_ (*)
// with &StructName (&Saiyan, in line 23) you can get the address in memory of the given variable value
func super_pointer(s *Saiyan){
  s.Power += 10000
}
