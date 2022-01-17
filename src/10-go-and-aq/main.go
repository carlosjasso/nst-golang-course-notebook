package main

import "fmt"

type BaseUnitOfTemperature string

const (
	Celciius BaseUnitOfTemperature = "C"
	Kelvin   BaseUnitOfTemperature = "K"
)

type Temperature float64

func (t Temperature) Format(b BaseUnitOfTemperature) string {
	return fmt.Sprintf("%.2f °%s", t, b)
}

func main() {
	var currTemp Temperature = 2.75
	fmt.Println("Current temp: ", currTemp.Format(Kelvin))
}
