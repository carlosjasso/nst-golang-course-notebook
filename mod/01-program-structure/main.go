/*
Each package must be provided with its own "main" package to provide the runtime an entrypoint
*/

package main

/*
Import all required packages
*/

import (
	"fmt" // "fmt" stands for "formatted I/O"
	"nst-golang-course/01-program-structure/models"
)

/*
Each main package must have its own "main" fuction to provide the runtime an entrypoint
*/

func main() {
	var planet = models.Planet{Name: "Krypton", Population: 1400000000} // Var declaration with initialization
	fmt.Printf("Hello from planet %s which has %d inhabitants.\n\n", planet.Name, planet.Population)

	var p models.Planet // variable declaration. All the struct fields are initialized with their default value
	fmt.Printf("Planet with default values: %v\n", p)
	p.Name = "madeUp"
	p.Population = 123
	p.SetAlias("planet alias")
	fmt.Printf("Planet with modified values: %v", p)
}
