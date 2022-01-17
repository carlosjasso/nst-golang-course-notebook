package main

import (
	"fmt"
	common "nst-golang-course/00-common"
	"nst-golang-course/09-functions-extras/models"
)

func main() {
	common.LabelMessage("User functions")
	c := models.Car{
		Model: "Somemodel",
		Make:  "Somemake",
	}

	u := models.User{
		Username:  "Someusername",
		Firstname: "Somename",
		Lastname:  "Somelastname",
		Salary:    100.0,
		Car:       &c,
	}

	fmt.Println("User: ", u.Info())
	newSalary, _ := u.IncreaseSalary(.5)
	fmt.Println("New salary: ", newSalary)
	fmt.Println("User car: ", u.Car.Info())
}
