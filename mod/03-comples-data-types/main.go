package main

import (
	"encoding/json"
	"fmt"
	common "nst-golang-course/00-common"
	types "nst-golang-course/03-comples-data-types/types"
	"strconv"
)

func main() {
	common.LabelMessage("Int arrays")
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	fmt.Println("Int array: ", numbers)

	common.LabelMessage("String arrays")
	// lazy init
	var names = [3]string{"name1", "name2", "name3"}
	fmt.Println("names: ", names)
	fmt.Println("Name@Index 1: ", names[1])

	// No hard sized array
	var elements = []int{4, 3, 2, 1, 0}
	common.LabelMessage("Lopped array iteration")
	for i := 0; i < len(elements); i++ {
		var message string = "Element @ index " + strconv.Itoa(i) + ": "
		fmt.Println(message, names)
	}

	// Array appending
	elements = append(elements, -1)
	elements = append([]int{5}, elements...)

	common.LabelMessage("Another lopped array iteration")
	for index, value := range elements {
		var message string = "Element @ index " + strconv.Itoa(index) + ": "
		fmt.Println(message, value)
	}

	common.LabelMessage("Maps!")
	ages := map[string]int{
		"name1": 10,
		"name2": 12,
		"name3": 11,
	}

	ages["name1"] = 13

	fmt.Println("Ages: ", ages)

	// Nested structs
	common.LabelMessage("Nested structs")
	var person = types.Person{
		Name: "Gofer",
		Age:  15,
		Car: types.Car{
			Model: "GoCar",
			Year:  2022,
		},
	}

	fmt.Println("Person: ", person)

	// JSON!
	common.LabelMessage("JSON!")
	// serialization
	data, err := json.Marshal(person)
	if err != nil {
		fmt.Println("!ERROR: \n\t", err)
	}
	fmt.Println("json bytes: ", data)

	// deserialization
	var deserialized types.Person

	err = json.Unmarshal(data, &deserialized)
	if err != nil {
		fmt.Println("!ERROR: \n\t", err)
	}
	fmt.Println("json: ", deserialized)
}
