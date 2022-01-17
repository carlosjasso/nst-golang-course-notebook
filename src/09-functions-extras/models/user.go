package models

import (
	"errors"
	"fmt"
)

var (
	ErrorWrongSalaryPercentage = errors.New("wrong salary increase percentage")
)

type User struct {
	Username  string
	Firstname string
	Lastname  string
	Salary    float64
	Car       *Car
}

func (u User) Info() string {
	return fmt.Sprintf("Username: %v\n\tFirstname: %v\n\tLastname: %v\n\tSalary: %v", u.Username, u.Firstname, u.Lastname, u.Salary)
}

func (u *User) IncreaseSalary(percentage float64) (float64, error) {
	if percentage > .5 {
		return 0, ErrorWrongSalaryPercentage
	}
	u.Salary += u.Salary * percentage
	return u.Salary, nil
}
