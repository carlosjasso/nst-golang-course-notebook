package main

import (
	"fmt"
	common "nst-golang-course/00-common"
	"nst-golang-course/05-functions/funcs"
	"nst-golang-course/05-functions/models"
)

func main() {
	common.LabelMessage("Functions!")

	// mean
	data := []float64{5, 6, 4, 7, 3, 8, 2, 9, 1}
	fmt.Println("Data: ", data)
	m := funcs.Mean(data...)
	fmt.Println("Mean: ", m)

	// multiple return
	_, lastname, _ := funcs.Person()
	fmt.Println("Lastname: ", lastname)

	// Recursion
	common.LabelMessage("recursion")
	post := models.Post{
		Text:     "#BLACK #BLACK #BLACK",
		Hasthags: nil,
		Posts: []models.Post{
			{
				Text:     "This is some #Random text",
				Hasthags: nil,
				Posts:    nil,
			},
			{
				Text:     "Totally #NotRandom text",
				Hasthags: nil,
				Posts:    nil,
			},
		},
	}

	fmt.Println("Post: ", post)
	post.ProcessHashtags()

	// Anonimous functions
	callback := funcs.Autoincrement()
	fmt.Println("current value of x: ", callback())
	fmt.Println("current value of x: ", callback())
	fmt.Println("current value of x: ", callback())
}
