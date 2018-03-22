// All go progams need a main package
package main

// Format package for printing and other stuff
import "fmt"

// All go programas have a main funciont which is the
// entrypoint of the application
func main() {
	// Declared variables need to have their type specified
	var a int
	// an already declared variable can have a value of the specified type
	// assigned to it
	a = 1

	// You can also let Go infer the type of the variable by using the
	// := shortcut
	e := 2

  // Multiple variables can be declared and have their value assigned(type is also inferred)
	var c, d = 3, 4
	// even with the := shortcut
	f, g := 5, 6

	fmt.Println(a, e, c, d, f, g)
}
