package models

import "fmt"

type Car struct {
	Model string
	Make  string
}

func (c Car) Info() string {
	return fmt.Sprintf("\n\tModel: %v\n\tMake: %v", c.Model, c.Make)
}
