package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"nst-golang-course/04-html-server/types"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		paginator := types.Paginator{
			Title: "Home",
			People: []types.Person{
				{
					Name:   "Name1",
					Age:    10,
					Active: true,
				},
				{
					Name:   "Name2",
					Age:    12,
					Active: false,
				},
				{
					Name:   "Name3",
					Age:    14,
					Active: true,
				},
			},
		}

		// Path relative to main.go location or execution location
		tmpl := template.Must(template.ParseFiles("./templates/index.gohtml"))

		if err := tmpl.Execute(w, paginator); err != nil {
			fmt.Println("!ERROR: ", err)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
